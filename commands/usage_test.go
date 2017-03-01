package commands_test

import (
	"bytes"
	"strings"

	"github.com/cloudfoundry/bosh-bootloader/commands"
	"github.com/cloudfoundry/bosh-bootloader/storage"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usage", func() {
	var (
		usage  commands.Usage
		stdout *bytes.Buffer
	)

	BeforeEach(func() {
		stdout = bytes.NewBuffer([]byte{})
		usage = commands.NewUsage(stdout)
	})

	Describe("Execute", func() {
		It("prints out the usage information", func() {
			err := usage.Execute([]string{}, storage.State{})
			Expect(err).NotTo(HaveOccurred())
			Expect(stdout.String()).To(Equal(strings.TrimLeft(`
Usage:
  bbl [GLOBAL OPTIONS] COMMAND [OPTIONS]

Global Options:
  --help      [-h]       Prints usage
  --state-dir            Directory containing bbl-state.json
  --debug                Prints debugging output

Commands:
  bosh-ca-cert           Prints BOSH director CA certificate
  bosh-deployment-vars   Prints required variables for BOSH deployment
  create-lbs             Attaches load balancer(s)
  delete-lbs             Deletes attached load balancer(s)
  destroy                Tears down BOSH director infrastructure
  director-address       Prints BOSH director address
  director-username      Prints BOSH director username
  director-password      Prints BOSH director password
  director-ca-cert       Prints BOSH director CA certificate
  env-id                 Prints environment ID
  print-env              Prints BOSH friendly environment variables
  help                   Prints usage
  lbs                    Prints attached load balancer(s)
  ssh-key                Prints SSH private key
  up                     Deploys BOSH director on AWS
  update-lbs             Updates load balancer(s)
  version                Prints version

  Use "bbl [command] --help" for more information about a command.
`, "\n")))
		})
	})

	Describe("PrintCommandUsage", func() {
		It("prints the usage for given command", func() {
			usage.PrintCommandUsage("my-command", "some message")
			Expect(stdout.String()).To(Equal(strings.TrimLeft(`Usage:
  bbl [GLOBAL OPTIONS] my-command [OPTIONS]

Global Options:
  --help      [-h]       Prints usage
  --state-dir            Directory containing bbl-state.json
  --debug                Prints debugging output

[my-command command options]
  some message
`, "\n")))
		})
	})
})
