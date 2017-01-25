package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cloudfoundry/bosh-bootloader/aws"
	"github.com/cloudfoundry/bosh-bootloader/aws/cloudformation"
	"github.com/cloudfoundry/bosh-bootloader/aws/ec2"
	"github.com/cloudfoundry/bosh-bootloader/aws/iam"
	"github.com/cloudfoundry/bosh-bootloader/bosh"
	"github.com/cloudfoundry/bosh-bootloader/storage"
)

const (
	UpCommand = "up"
)

type keyPairSynchronizer interface {
	Sync(keypair ec2.KeyPair) (ec2.KeyPair, error)
}

type infrastructureManager interface {
	Create(keyPairName string, numberOfAZs int, stackName, lbType, lbCertificateARN, envID string) (cloudformation.Stack, error)
	Update(keyPairName string, numberOfAZs int, stackName, lbType, lbCertificateARN, envID string) (cloudformation.Stack, error)
	Exists(stackName string) (bool, error)
	Delete(stackName string) error
	Describe(stackName string) (cloudformation.Stack, error)
}

type availabilityZoneRetriever interface {
	Retrieve(region string) ([]string, error)
}

type credentialValidator interface {
	ValidateAWS() error
	ValidateGCP() error
}

type logger interface {
	Step(string, ...interface{})
	Println(string)
	Prompt(string)
}

type stateStore interface {
	Set(state storage.State) error
}

type certificateDescriber interface {
	Describe(certificateName string) (iam.Certificate, error)
}

type configProvider interface {
	SetConfig(config aws.Config)
}

type AWSUp struct {
	credentialValidator       credentialValidator
	infrastructureManager     infrastructureManager
	keyPairSynchronizer       keyPairSynchronizer
	boshDeployer              boshDeployer
	boshCloudConfigurator     boshCloudConfigurator
	availabilityZoneRetriever availabilityZoneRetriever
	certificateDescriber      certificateDescriber
	cloudConfigManager        cloudConfigManager
	boshClientProvider        boshClientProvider
	envIDGenerator            envIDGenerator
	stateStore                stateStore
	configProvider            configProvider
}

type AWSUpConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

func NewAWSUp(
	credentialValidator credentialValidator, infrastructureManager infrastructureManager,
	keyPairSynchronizer keyPairSynchronizer, boshDeployer boshDeployer,
	boshCloudConfigurator boshCloudConfigurator, availabilityZoneRetriever availabilityZoneRetriever,
	certificateDescriber certificateDescriber, cloudConfigManager cloudConfigManager,
	boshClientProvider boshClientProvider, stateStore stateStore,
	configProvider configProvider) AWSUp {

	return AWSUp{
		credentialValidator:       credentialValidator,
		infrastructureManager:     infrastructureManager,
		keyPairSynchronizer:       keyPairSynchronizer,
		boshDeployer:              boshDeployer,
		boshCloudConfigurator:     boshCloudConfigurator,
		availabilityZoneRetriever: availabilityZoneRetriever,
		certificateDescriber:      certificateDescriber,
		cloudConfigManager:        cloudConfigManager,
		boshClientProvider:        boshClientProvider,
		stateStore:                stateStore,
		configProvider:            configProvider,
	}
}

func (u AWSUp) Execute(config AWSUpConfig, state storage.State) error {
	state.IAAS = "aws"

	if u.awsCredentialsPresent(config) {
		state.AWS.AccessKeyID = config.AccessKeyID
		state.AWS.SecretAccessKey = config.SecretAccessKey
		state.AWS.Region = config.Region
		if err := u.stateStore.Set(state); err != nil {
			return err
		}
		u.configProvider.SetConfig(aws.Config{
			AccessKeyID:     config.AccessKeyID,
			SecretAccessKey: config.SecretAccessKey,
			Region:          config.Region,
		})
	} else if u.awsCredentialsNotPresent(config) {
		err := u.credentialValidator.ValidateAWS()
		if err != nil {
			return err
		}
	} else {
		return u.awsMissingCredentials(config)
	}

	err := u.checkForFastFails(state)
	if err != nil {
		return err
	}

	if state.KeyPair.Name == "" {
		state.KeyPair.Name = fmt.Sprintf("keypair-%s", state.EnvID)
	}

	if err := u.stateStore.Set(state); err != nil {
		return err
	}

	keyPair, err := u.keyPairSynchronizer.Sync(ec2.KeyPair{
		Name:       state.KeyPair.Name,
		PublicKey:  state.KeyPair.PublicKey,
		PrivateKey: state.KeyPair.PrivateKey,
	})
	if err != nil {
		return err
	}

	state.KeyPair.PublicKey = keyPair.PublicKey
	state.KeyPair.PrivateKey = keyPair.PrivateKey

	if err := u.stateStore.Set(state); err != nil {
		return err
	}

	availabilityZones, err := u.availabilityZoneRetriever.Retrieve(state.AWS.Region)
	if err != nil {
		return err
	}

	if state.Stack.Name == "" {
		state.Stack.Name = fmt.Sprintf("stack-%s", strings.Replace(state.EnvID, ":", "-", -1))

		if err := u.stateStore.Set(state); err != nil {
			return err
		}
	}

	var certificateARN string
	if lbExists(state.Stack.LBType) {
		certificate, err := u.certificateDescriber.Describe(state.Stack.CertificateName)
		if err != nil {
			return err
		}
		certificateARN = certificate.ARN
	}

	stack, err := u.infrastructureManager.Create(state.KeyPair.Name, len(availabilityZones), state.Stack.Name, state.Stack.LBType, certificateARN, state.EnvID)
	if err != nil {
		return err
	}

	deployInput := bosh.DeployInput{
		IAAS:                  "aws",
		DirectorName:          fmt.Sprintf("bosh-%s", state.EnvID),
		AZ:                    stack.Outputs["BOSHSubnetAZ"],
		AccessKeyID:           stack.Outputs["BOSHUserAccessKey"],
		SecretAccessKey:       stack.Outputs["BOSHUserSecretAccessKey"],
		Region:                state.AWS.Region,
		DefaultKeyName:        state.KeyPair.Name,
		DefaultSecurityGroups: []string{stack.Outputs["BOSHSecurityGroup"]},
		SubnetID:              stack.Outputs["BOSHSubnet"],
		ExternalIP:            stack.Outputs["BOSHEIP"],
		PrivateKey:            state.KeyPair.PrivateKey,
		BOSHState:             state.BOSH.State,
		Variables:             state.BOSH.Variables,
	}

	deployOutput, err := u.boshDeployer.Deploy(deployInput)
	if err != nil {
		return err
	}

	directorOutputs := getDirectorOutputs(deployOutput.Variables)

	variablesYAMLContents, err := marshal(deployOutput.Variables)
	if err != nil {
		return err
	}

	variablesYAML := string(variablesYAMLContents)
	state.BOSH = storage.BOSH{
		DirectorName:           deployInput.DirectorName,
		DirectorAddress:        stack.Outputs["BOSHURL"],
		DirectorUsername:       DIRECTOR_USERNAME,
		DirectorPassword:       directorOutputs.directorPassword,
		DirectorSSLCA:          directorOutputs.directorSSLCA,
		DirectorSSLCertificate: directorOutputs.directorSSLCertificate,
		DirectorSSLPrivateKey:  directorOutputs.directorSSLPrivateKey,
		Variables:              variablesYAML,
		State:                  deployOutput.BOSHState,
	}

	err = u.stateStore.Set(state)
	if err != nil {
		return err
	}

	boshClient := u.boshClientProvider.Client(state.BOSH.DirectorAddress, state.BOSH.DirectorUsername,
		state.BOSH.DirectorPassword)

	cloudConfigInput := u.boshCloudConfigurator.Configure(stack, availabilityZones)

	err = u.cloudConfigManager.Update(cloudConfigInput, boshClient)
	if err != nil {
		return err
	}

	err = u.stateStore.Set(state)
	if err != nil {
		return err
	}

	return nil
}

func (u AWSUp) checkForFastFails(state storage.State) error {
	stackExists, err := u.infrastructureManager.Exists(state.Stack.Name)
	if err != nil {
		return err
	}

	if !state.BOSH.IsEmpty() && !stackExists {
		return fmt.Errorf(
			"Found BOSH data in state directory, but Cloud Formation stack %q cannot be found "+
				"for region %q and given AWS credentials. bbl cannot safely proceed. Open an issue on GitHub at "+
				"https://github.com/cloudfoundry/bosh-bootloader/issues/new if you need assistance.",
			state.Stack.Name, state.AWS.Region)
	}

	return nil
}

func (AWSUp) awsCredentialsPresent(config AWSUpConfig) bool {
	return config.AccessKeyID != "" && config.SecretAccessKey != "" && config.Region != ""
}

func (AWSUp) awsCredentialsNotPresent(config AWSUpConfig) bool {
	return config.AccessKeyID == "" && config.SecretAccessKey == "" && config.Region == ""
}

func (AWSUp) awsMissingCredentials(config AWSUpConfig) error {
	switch {
	case config.AccessKeyID == "":
		return errors.New("AWS access key ID must be provided")
	case config.SecretAccessKey == "":
		return errors.New("AWS secret access key must be provided")
	case config.Region == "":
		return errors.New("AWS region must be provided")
	}

	return nil
}
