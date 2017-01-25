// Code generated by go-bindata.
// sources:
// vendor/github.com/cloudfoundry/bosh-deployment/bosh.yml
// vendor/github.com/cloudfoundry/bosh-deployment/external-ip-not-recommended.yml
// vendor/github.com/cloudfoundry/bosh-deployment/external-ip-with-registry-not-recommended.yml
// vendor/github.com/cloudfoundry/bosh-deployment/gcp/cpi.yml
// vendor/github.com/cloudfoundry/bosh-deployment/aws/cpi.yml
// DO NOT EDIT!

package bosh

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _vendorGithubComCloudfoundryBoshDeploymentBoshYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x0c\x52\x20\xf1\x06\x91\x6c\xc9\xf6\xda\x21\x50\xb4\x45\x2f\x2d\x90\x4b\x81\xf4\xd2\x20\x20\x28\x72\x2c\x31\x2b\x91\x02\x49\x69\xeb\x5d\xec\x7f\x2f\xa8\x2f\xcb\xb2\xbd\x5d\xec\xc1\xeb\x37\x8f\x8f\xc3\x99\xc7\x31\xc3\x30\x0c\x14\x2b\x91\x40\xaa\x6d\x1e\x04\x06\x0b\x64\x16\x2d\x09\x42\x98\xe0\x00\x0d\x1a\x2b\xb5\x22\x90\xdc\xaf\xa2\x75\x00\x50\x9b\x82\x40\xee\x5c\x65\xc9\x72\x69\xd7\x11\x2b\xd9\x93\x56\xec\xd1\x46\x5c\x97\x4b\xbf\x2a\xe4\xba\xac\x64\x81\x22\xec\x55\x43\xc7\x4c\xca\x8a\xc2\x2e\x07\xa0\xa5\xb5\x8a\xa1\x56\x61\x9d\xd6\xca\xd5\xa1\x33\xb5\x75\xc7\xd0\x3a\x2c\x39\x16\x45\xb8\x5e\xc7\x49\x14\x6f\xc3\x64\x15\xef\x56\x71\xbc\x59\xad\xb6\xf1\x66\x15\xb9\xec\xe9\x97\x3e\xab\x3f\xc5\xcf\x4f\xc6\x6c\xcb\x7f\xfe\xb2\xbf\x7d\xfd\xf2\x45\x7c\xb5\x59\x89\xf9\x3a\xff\xe3\xc7\xdf\xa8\xcd\x97\x5d\x93\xfd\x1e\x00\xd8\x9c\xc5\x04\xd6\x87\xe4\xf0\x79\x13\x27\xbb\x94\x6d\x30\x4d\xd3\x1d\x8f\x45\xcc\x93\xed\xe7\x34\xd9\xe0\x9a\xf3\x3d\x5f\xef\x31\x41\xf4\xb5\xb0\xba\x36\x1c\x69\xa5\x75\x31\xa9\x48\x53\xda\x00\x40\xa1\x7b\xd4\xe6\x81\x80\xc0\x03\xab\x0b\x17\x00\xa0\x6a\x48\x00\x00\x6d\xcd\xba\xff\x00\x2a\x66\xed\xa3\x36\x82\xc0\x87\x8f\x1f\x82\x40\x48\xfb\x30\x17\xf4\x98\x97\x6c\x63\x56\x3e\x21\x81\x75\x42\x77\xf7\xfb\x20\xe8\x77\x99\x92\xc7\xed\xdc\xb1\x42\x02\x25\x53\x35\x2b\xfc\xf1\xea\x54\xa1\xb3\x7e\xdf\x10\x0c\x53\x19\x12\x58\x2c\xa4\x72\x68\x14\x2b\x28\x97\xc2\xdc\xdd\xb5\x49\x65\xcc\xe1\x23\x3b\x9e\x85\xb3\xc7\x3e\x68\x1d\x73\x92\x13\xf8\x36\x09\xca\xea\xee\xee\x7b\x1b\x15\xca\x12\xf8\xb6\x8f\xda\xbf\xef\x41\x20\x95\x75\x4c\x71\xa4\x99\xd1\x75\x75\x69\x9b\x21\x6e\x09\xc4\x01\xc0\x0f\x9d\xf6\x09\x3e\x77\x34\xc5\x9c\xfd\x04\xbd\x1d\xba\x55\x2f\xd3\x78\xa5\xad\xcb\x0c\xda\xf0\x73\xb4\x79\x8d\x97\x16\x3a\xb5\x4e\x1b\x7c\x8d\x24\xa4\x41\xee\xb4\x79\x8d\x93\x23\x2b\x5c\x4e\x4b\xad\xe4\x75\xe6\x99\x2b\x06\x33\x54\xde\x87\xd6\xa1\x72\x74\x6c\xf0\xa9\xaf\xa7\x26\xfa\x8d\xe6\x6d\x1c\x4a\x4e\x65\x65\xaf\x97\xbd\x32\xba\x42\xe3\x24\xda\xce\x54\xbe\x68\x83\xbd\x98\x10\x06\xad\x2f\x6f\xb2\x8b\x56\xd1\x2a\x8a\xfb\x40\x6d\xd1\x74\xf5\xbd\x30\xe2\x62\xe1\x61\x3a\x00\x7d\xe3\x87\x52\x13\x78\x2f\xd2\x7e\x4d\xd1\x1e\x8a\xde\xdc\x24\xd7\xd6\xdd\xda\x7a\xd0\xbb\xb2\xfd\x10\x9a\xa7\x00\x20\x98\x63\xe9\x58\xef\xf1\x8c\xac\x72\x17\x92\x63\xcb\x2f\x4a\x31\x2b\xe1\x90\x80\x36\x8e\x40\xb2\x4d\xb6\xab\x01\x31\xba\x91\xc2\x0b\x0b\xd6\x0c\x09\xf4\x1e\x19\x44\x87\xd3\x0c\xf8\x08\x4f\xcf\x33\xa6\x42\x07\xda\xe5\xc9\x58\x86\xca\xcd\x55\x5b\xf0\x7f\x24\x5b\xce\x5c\x6f\x9e\xe6\xcd\x0e\x75\x6e\x5b\x2c\xc6\xc4\x3c\x70\x2a\x77\x4a\xe0\xe3\xd8\xed\x43\x51\xdb\x9c\x32\x53\x11\x70\xa6\xc6\x1e\x45\xc5\xd2\xc2\xdb\xdd\x3a\x2a\xb0\x2a\xf4\xf1\x2c\x9c\xa1\x42\xc3\x1c\xd2\xa6\x1c\x93\xb4\xd7\x04\x04\x0a\xc9\x99\x43\x41\xbd\xdf\x6b\x4b\xfd\x9d\xf0\x45\x98\x50\x3b\xc8\x12\xd8\x0c\x6b\x1b\x54\x27\xb3\xfb\xdb\xc7\xb5\x11\xb4\x87\xa7\x4b\xad\x2d\x4e\xb4\x07\x3c\x9e\x1d\xda\xda\x22\xaa\x8c\x6c\x7c\x9e\x0f\x78\x1c\xcf\x0f\xc0\xd1\xbb\x62\x46\xf5\xa0\x3c\xb4\xd9\x8e\x54\xdf\x30\x5a\x32\xc5\x32\x2c\xcf\x3a\x79\x72\x51\xa1\x79\x3b\x89\xfb\xbb\xe3\xbf\x9d\x68\x9d\x82\x9d\x02\xc3\x2c\x60\xa2\x94\x6a\x82\x9f\x3b\xa1\x8d\x5e\xfa\x69\x2a\x90\x97\x37\x57\xe7\xe5\x7c\x69\x5e\x92\x99\xd7\x29\xe3\x5c\xd7\x97\xee\x9c\xc8\xbe\x2e\xd9\x16\x92\xd1\xeb\xb5\x64\x23\xc9\xa0\xad\x4d\x1f\xea\x4c\x21\x26\x2d\x54\xae\x22\xf0\x5e\xb9\x2a\xe8\x8e\xe6\x64\x89\x71\x94\x69\x9d\x15\xe8\xdf\x14\x13\x38\xb9\x0e\xaf\xaf\xc3\x9b\x39\x7c\x76\x15\xcb\xb4\xb6\xdd\x9c\x24\xcb\x65\xfb\x71\x31\x1e\x7f\x9d\x0d\x14\xb2\x49\x92\x24\x08\x78\xa1\x6b\x41\xc7\xf6\x07\x83\xd6\xf0\x28\x6a\xbf\x2d\x16\xfe\x83\xa6\x5a\x3b\xeb\x0c\xab\x5e\x93\xbd\xdf\xdf\xef\xaf\x8c\xfb\x2e\x5d\x78\xee\xd4\xdf\xbd\x59\x7e\xd5\x0e\x82\x55\x2b\xfb\xee\x65\x36\x33\xe1\x79\xe6\xdb\x4f\x50\x31\x97\x13\x58\x36\xcc\x2c\x1b\xce\xaa\x65\x29\xb9\xd1\xd4\xcf\xe1\xa5\x1f\xcb\x4b\xce\x78\x8e\x2f\xa7\x5e\x7d\xf4\xbd\x0a\x1a\x66\xa4\x6f\xe5\xe4\xb7\xff\xdc\xb1\xe3\x4b\x65\x04\xc6\x37\xc2\xed\xb1\xf9\x96\x45\xe7\x83\xf1\xf6\x8a\x89\x5f\x6f\x93\x6e\x94\xf1\xf6\x82\x33\x8f\xdc\xa6\x5d\xfc\xd4\xdd\xa6\xf6\x4f\x02\xca\xd9\xc8\x99\x0c\xa2\x00\x40\x57\x4e\x6a\xd5\xbb\x82\xeb\xb2\xd4\x8a\x76\x4b\x39\x83\x9f\xc0\x69\xa1\xe1\x20\xff\x9d\xbc\x2b\x4f\xf7\xf0\x4d\x92\x6c\x96\xc4\x6c\x97\x6b\xbf\xab\xac\x68\x11\x27\x1b\x6c\x59\x57\x9f\x30\xff\x05\x00\x00\xff\xff\x52\xa3\x94\x77\x66\x0c\x00\x00")

func vendorGithubComCloudfoundryBoshDeploymentBoshYmlBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComCloudfoundryBoshDeploymentBoshYml,
		"vendor/github.com/cloudfoundry/bosh-deployment/bosh.yml",
	)
}

func vendorGithubComCloudfoundryBoshDeploymentBoshYml() (*asset, error) {
	bytes, err := vendorGithubComCloudfoundryBoshDeploymentBoshYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/cloudfoundry/bosh-deployment/bosh.yml", size: 3174, mode: os.FileMode(420), modTime: time.Unix(1484943925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComCloudfoundryBoshDeploymentExternalIpNotRecommendedYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x93\x41\x6b\x1c\x31\x0c\x85\xef\xf3\x2b\x04\xb9\xec\x42\xb7\xee\x29\x04\x43\x49\x0b\x85\xd2\x4b\x0e\xbd\xe4\x10\x8a\xd1\xd8\x4a\xc6\xd4\x6b\x19\x49\xe3\xcd\xfe\xfb\xb2\x93\x49\xd2\x86\x36\x85\xf6\x64\x90\xed\xa7\xf7\x3e\xd9\x67\x70\xfd\xf1\xeb\xd5\x97\xab\xcf\x1e\xae\x09\x12\x43\x65\x03\xa1\xc8\xfb\x3d\xd5\x04\xc6\x40\xf7\x8d\x95\xe0\x53\x16\x8a\xc6\x02\x58\x13\x60\x3d\x02\xdb\x44\x02\x4a\xd2\x73\x24\x05\xee\x24\xd0\xe6\xb1\xe4\x08\x98\x92\x90\x2a\xe9\xdb\x61\xd8\x81\x1d\x1b\x79\x10\x6a\x05\x23\x0d\x00\x0d\x6d\xf2\xe0\x2a\xd9\x81\xe5\xbb\xba\xdd\x00\xd0\xb1\xcc\xe4\x07\x00\x80\x8a\x7b\xf2\xab\xd2\x52\x78\xb8\xdf\x73\x7b\x45\x2c\x57\x35\xac\x91\xc2\x9d\xf0\xdc\xd4\x9d\x44\xde\x8f\xac\xd3\x73\x9b\x77\x2e\xd1\x2d\xce\xc5\x2e\x9f\xfa\xc1\x4d\xaa\xfa\x06\xee\xd0\xe8\x80\xc7\x6f\xff\xa7\xff\xb7\x18\x6a\x68\x39\x86\xdc\xd4\xc3\xcd\x66\x43\xf7\x46\x52\xb1\x84\xdc\xb6\xdb\x7f\x6b\xdd\x84\x1b\x89\x65\x52\x97\xd6\xe9\x3c\x66\x0c\xaa\x53\xe0\x66\x99\xab\x5e\xfe\x6a\x6c\x4d\x1b\x26\x56\xf3\xf0\xc2\xc8\x30\x9c\x81\x71\x62\xd0\x89\xe7\x92\x96\xd7\x80\x31\x92\x2a\x54\xae\xbb\x44\xb7\xb9\x52\x82\x8e\xa2\x7f\x76\x1c\x0b\xcf\x29\x34\xe1\x9e\x13\x89\xdb\x8f\xb3\x3e\x23\x9f\xcc\x9a\x7a\xb7\x54\xfd\x66\x73\x5a\xc2\xc8\x6c\x6a\x82\x2d\x34\x54\x3d\xb0\xa4\xed\xf6\xc3\x0b\x67\xfe\xfc\xe2\xfc\xe2\x15\x4c\x1d\x25\xe3\x58\x68\x05\xf4\x08\x24\xa8\x16\xb7\x82\x70\x58\x16\x41\xcb\x9d\xc2\xe9\xd4\xcf\x43\xfb\x1d\x09\x3d\x64\x8b\x13\xcc\x88\x0f\x1f\xe1\x69\x7b\xf8\x11\x00\x00\xff\xff\xdb\x7f\x2e\xb7\x37\x03\x00\x00")

func vendorGithubComCloudfoundryBoshDeploymentExternalIpNotRecommendedYmlBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComCloudfoundryBoshDeploymentExternalIpNotRecommendedYml,
		"vendor/github.com/cloudfoundry/bosh-deployment/external-ip-not-recommended.yml",
	)
}

func vendorGithubComCloudfoundryBoshDeploymentExternalIpNotRecommendedYml() (*asset, error) {
	bytes, err := vendorGithubComCloudfoundryBoshDeploymentExternalIpNotRecommendedYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/cloudfoundry/bosh-deployment/external-ip-not-recommended.yml", size: 823, mode: os.FileMode(420), modTime: time.Unix(1484943925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComCloudfoundryBoshDeploymentExternalIpWithRegistryNotRecommendedYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x93\xcd\x6a\xdc\x30\x10\xc7\xef\x7e\x8a\x81\x5c\x76\xa1\x5b\xf5\x14\x82\xa0\xa4\x85\x42\xe9\x25\x87\x5e\x72\x08\x45\x8c\xa5\x49\x2c\xaa\xd5\x88\x99\xb1\x37\xfb\xf6\xc5\x1b\x27\xdb\x6e\x9b\x2d\xb4\x27\xc3\xd8\xfe\x7f\xfc\x34\xba\x80\xdb\x8f\x5f\x6f\xbe\xdc\x7c\xf6\x70\x4b\x90\x18\x2a\x1b\x08\x45\xde\x6e\xa9\x26\x30\x06\x7a\x6c\xac\x04\x9f\xb2\x50\x34\x16\xc0\x9a\x00\xeb\x1e\xd8\x06\x12\x50\x92\x29\x47\x52\xe0\x89\x04\xda\xd8\x97\x1c\x01\x53\x12\x52\x25\x7d\xdb\x75\x1b\xb0\x7d\x23\x0f\x42\xad\x60\xa4\x0e\xa0\xa1\x0d\x1e\x5c\x25\xdb\xb1\x7c\x57\xb7\xe9\x00\x26\x2c\x23\xf9\x0e\x00\xa0\xe2\x96\xfc\xa2\x74\x18\x3c\xfd\x3f\xe5\x76\x46\x2c\x57\x35\xac\x91\xc2\x83\xf0\xd8\xd4\xcd\x22\xef\x7b\xd6\xe1\x68\xf3\xce\x25\xba\xc7\xb1\xd8\xf5\x8b\x1f\xdc\xa5\xaa\x6f\xe0\x01\x8d\x76\xb8\xff\xf6\x7f\xfa\x7f\xab\xa1\x86\x96\x63\xc8\x4d\x3d\xdc\xad\x56\xf4\x68\x24\x15\x4b\xc8\x6d\xbd\xfe\x37\xeb\x26\xdc\x48\x2c\x93\xba\xb4\x9c\xce\x73\xc7\xa0\x3a\x04\x6e\x96\xb9\xea\xf5\xaf\xc1\x96\xb6\x61\x60\x35\x0f\x27\x41\xba\xee\x02\x8c\x13\x83\x0e\x3c\x96\x74\xd8\x06\x8c\x91\x54\xa1\x72\xdd\x24\xba\xcf\x95\x12\x4c\x28\xfa\x7a\xe2\x58\x78\x4c\xa1\x09\x4f\x39\x91\xb8\x6d\x3f\xea\x11\xf9\x60\xd6\xd4\xbb\xc3\xd4\xaf\x56\xf3\x23\xf4\xcc\xa6\x26\xd8\x42\x43\xd5\x1d\x4b\x5a\xaf\x3f\x9c\x24\xf3\x97\x57\x97\x57\x67\x30\x9d\x98\xce\xfd\x6d\xac\x95\x8a\x9b\x8b\x1e\xfd\x7f\x2b\xfc\xaa\xe2\x84\x92\xb1\x2f\xb4\x20\x7f\x46\x1c\x54\x8b\x5b\xd0\x3a\x2c\x07\x2d\xcb\x13\x85\xf9\xab\x9f\xd7\xe0\x4f\x6c\x75\x97\x2d\x0e\x30\x22\x3e\x5d\xad\x97\xd7\xdd\x8f\x00\x00\x00\xff\xff\xd0\xc7\xbe\x5c\x89\x03\x00\x00")

func vendorGithubComCloudfoundryBoshDeploymentExternalIpWithRegistryNotRecommendedYmlBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComCloudfoundryBoshDeploymentExternalIpWithRegistryNotRecommendedYml,
		"vendor/github.com/cloudfoundry/bosh-deployment/external-ip-with-registry-not-recommended.yml",
	)
}

func vendorGithubComCloudfoundryBoshDeploymentExternalIpWithRegistryNotRecommendedYml() (*asset, error) {
	bytes, err := vendorGithubComCloudfoundryBoshDeploymentExternalIpWithRegistryNotRecommendedYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/cloudfoundry/bosh-deployment/external-ip-with-registry-not-recommended.yml", size: 905, mode: os.FileMode(420), modTime: time.Unix(1484943925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComCloudfoundryBoshDeploymentGcpCpiYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\x4d\x6f\xe3\x36\x10\xbd\xeb\x57\x10\x08\x50\xd8\x01\x64\xca\x96\x65\xd9\x02\x02\xa3\xc8\xa1\xe8\xa5\xc8\xa1\x3d\x15\x05\x41\x91\x23\x99\x36\x45\x12\xfc\x70\xd7\x59\xec\x7f\x5f\xd0\x96\x13\xd9\x41\xbc\xbb\xc9\x89\x22\x67\xe6\xcd\x7b\xf3\x34\x69\x9a\x26\x29\xf2\x07\x03\x15\xb2\x60\x24\x65\x90\x20\x64\xa8\xdf\x54\x08\x5b\x90\x40\x1d\x38\x9c\x26\x08\xed\xa9\x0c\x50\x25\x08\x21\xa4\x68\x07\x15\xaa\xb5\xdb\xa4\xad\xd6\xad\x84\x94\x19\x71\x8c\xec\xc1\x3a\xa1\x55\x85\x66\xc5\x64\x31\x99\x1e\xdf\x82\x95\x15\xda\x78\x6f\x5c\x85\x71\x2c\x9a\x08\x8d\x39\x6e\x85\xdf\x84\x7a\xc2\x74\x87\x99\xd4\x81\x37\x3a\x28\x6e\x0f\xa9\x50\x2c\xd4\xd4\x6b\x8b\xaf\x1a\xa4\x3d\x9d\xf5\xfe\x61\x80\xee\x36\x74\x5a\xa1\x86\x2d\x59\x99\xcd\x0a\x96\xd7\x34\xa7\xd0\x2c\x16\x90\x65\x05\x9d\xd7\xbc\x29\x1b\x9e\xf3\x02\x56\xab\x72\x55\xce\x93\x5b\x62\x9d\x0e\x96\x01\x31\x5a\x4b\x87\xa3\xc6\x87\x7d\xe7\xb0\xf3\xd0\x31\x90\x72\x7d\x39\x83\x77\x54\x9d\xb3\xdd\x05\xfb\xdd\xbe\x4b\x43\x1d\x94\x0f\xa9\xb7\xc1\xf9\x43\xda\x6a\x42\x5b\x50\x7e\xbd\x7f\xc8\xf3\xe9\x6c\x32\x2d\x06\x6a\x72\xca\x72\x80\x65\x5e\x16\x59\x53\x16\x35\x2f\xe7\xb0\xe4\x39\xe4\x94\xaf\xca\x65\xb6\xe4\xf5\x2c\x67\x79\x56\xd3\x24\xb9\x43\x8f\x5a\x35\xa2\x0d\x16\x90\x13\xcf\xe0\x7e\x5d\xdf\x71\xf8\xc4\x58\x6d\xc0\x7a\x01\xee\x4a\xe7\xb3\x56\x50\xa1\xd1\x28\x9e\xe3\xf1\xf1\xa9\xa3\x6c\x23\x14\x90\x53\x23\x35\x4d\x9d\xa7\x8a\x53\xcb\xd3\x93\x25\x56\x6b\x4f\xb8\x70\x3b\x12\x29\x91\xb6\xae\xd0\x3c\xbb\x8a\x9c\x6a\x0d\x7f\xa9\xbd\xe1\xcc\xb1\x60\xc0\x3a\xde\x6f\xf1\x46\x5f\xdf\xa0\x7f\xbb\x01\xaf\xc0\xff\xaf\xed\xee\x0c\x0e\x0d\x0d\xd2\x63\x17\x6a\x05\xde\xe1\xec\x47\x03\xea\xcb\xc9\x69\x29\x46\xa3\xfe\xde\xcf\xea\x04\x73\x91\xf0\xfa\xd4\xe7\x80\xd9\x40\x07\x96\x4a\x02\x5f\x3c\x58\x45\x25\x11\xa6\x42\x0d\x95\x0e\x8e\x09\x9e\xb6\x2e\x56\xc6\x73\x3c\x8e\xae\xff\xce\x39\x7a\x7c\xfa\x13\x6d\x75\xfd\xbe\x30\xa1\xa2\x7c\x06\xa4\xb5\x3a\x98\x5e\x5f\xfc\x2d\xf1\x56\xd7\xc3\xa5\x46\xbf\x31\x23\x48\xc4\x7a\xdd\xee\xd3\x9f\x4b\xce\x8b\xdd\xef\xde\xdb\xad\xff\x48\xfb\xd7\x59\x62\x2e\x2c\xb0\xb8\xec\x3d\x83\x81\x8b\x03\x06\xef\x37\x79\x31\x67\x2f\x38\x58\xec\xa1\x33\x92\x7a\x18\xc0\xdc\x9f\xb5\x7d\x92\xe9\x89\xce\xfa\x6a\x68\x4c\xab\xe6\x38\x20\x63\xf5\x16\x98\x8f\x36\xf5\x9f\x44\xf0\xde\xe0\xad\xd3\x8a\xec\xe0\x10\x83\x2d\x33\x84\x59\xe0\xa0\xbc\xa0\xd2\x91\x18\x8b\x9e\xfe\xac\xc2\x5b\x8c\xee\x5f\x18\x25\x77\xe8\x1f\x07\xe8\x8f\xc7\x27\xf4\xd7\xdf\x4f\x9f\x54\xae\xbc\x19\xa8\x56\xde\xa0\x7f\xa7\x8b\xd5\x64\x56\xcc\x27\xfd\xf9\xdf\x87\xf8\x5f\xe0\xde\xc7\xdb\xf7\x00\x00\x00\xff\xff\xfa\x11\x76\x52\x8e\x06\x00\x00")

func vendorGithubComCloudfoundryBoshDeploymentGcpCpiYmlBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComCloudfoundryBoshDeploymentGcpCpiYml,
		"vendor/github.com/cloudfoundry/bosh-deployment/gcp/cpi.yml",
	)
}

func vendorGithubComCloudfoundryBoshDeploymentGcpCpiYml() (*asset, error) {
	bytes, err := vendorGithubComCloudfoundryBoshDeploymentGcpCpiYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/cloudfoundry/bosh-deployment/gcp/cpi.yml", size: 1678, mode: os.FileMode(420), modTime: time.Unix(1484943925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComCloudfoundryBoshDeploymentAwsCpiYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x4d\x8f\xdb\x38\x0c\xbd\xfb\x57\x08\x18\xa0\x98\x14\xeb\x58\xf9\x1a\x27\x06\x06\x83\xa2\xd8\xc3\xde\x16\xd8\xc3\x1e\x0d\xda\x62\x12\x75\x1c\x49\x10\x25\xa7\x99\xa2\xff\x7d\x21\x7f\x24\x4e\x5a\x67\xb7\x98\xbd\x59\xa4\xf4\xc8\x47\xf2\xd1\x71\x1c\x47\x31\x73\x27\x83\x19\xb3\x68\x2a\x28\x31\x62\xcc\x80\xdb\x67\x2c\xb1\x58\x21\x10\x52\x12\x47\x8c\xd5\x50\x79\xcc\x22\xc6\x18\x53\x70\xc0\x8c\x15\x9a\xf6\x31\x1c\x29\x2e\x8d\x6c\xcc\x35\x5a\x92\x5a\x65\xec\x89\x37\x67\x6f\xab\x8c\xed\x9d\x33\x94\x25\x49\xb8\x3d\x95\x3a\x11\xc9\x4e\xba\xbd\x2f\xa6\xa5\x3e\x24\x65\xa5\xbd\xd8\x6a\xaf\x84\x3d\xc5\x52\x95\xbe\x00\xa7\x6d\x32\x44\x8e\xbb\x24\x5e\xea\xe7\x0e\x96\xf6\x30\xcb\xd8\x1a\x97\x1c\x36\xdb\xed\x7a\x33\x9f\xf3\x25\xe7\xe9\x7a\xbd\xe1\x8b\x74\xcb\x37\x4b\x98\x15\x5c\xcc\x17\x69\x9a\xf2\xd5\x3a\xba\x47\x8f\xb4\xb7\x25\xe6\x46\xeb\x8a\x92\xc0\xea\xb9\x3e\x50\x42\x0e\x0f\x25\x56\xd5\xcb\x35\xeb\x11\x3a\xfd\x6d\xba\xa4\xfd\x15\x55\xbc\xaf\x0f\xb1\x2f\xbc\x72\x3e\x76\xd6\x93\x3b\xc5\x3b\x9d\xc3\x0e\x95\x7b\xa9\x9f\x17\x8b\xd9\x7c\x3a\x5b\x0d\xe8\x6c\x57\x98\xae\xb7\x4f\xb3\x62\xcd\xe7\x4f\x8b\x34\x5d\x70\x2e\x8a\xd5\x72\x89\x1b\xb1\x28\x81\x17\xeb\xcd\x6a\xf6\x04\xdb\x45\x14\x3d\xb0\xcf\x5a\x6d\xe5\xce\x5b\x64\x9f\xfe\xfe\x8b\x91\x7c\x43\xfa\x75\x92\x4d\xe9\x73\x63\xb5\x41\xeb\x24\xd2\x0d\x59\xa9\xc8\x81\x2a\x31\x6f\x61\x0f\x8b\xe9\xd7\x0a\xec\x0e\x1b\x27\x9a\x3d\x1e\xd0\x42\x95\x0b\x49\xaf\x19\xfb\x16\x72\xc8\xd8\x7c\x95\x73\xce\x7f\xeb\x32\xd9\x99\xf9\xf7\xe6\x36\xd4\x20\x2b\x28\x64\x25\xdd\x29\x7f\xd3\x0a\x33\xf6\xf8\x08\x6f\x93\xc9\x9d\xd6\x04\xdc\x61\xc6\xe1\x7c\x2f\x67\xf6\x6d\x10\x74\x1c\x56\xa1\x3b\x6a\xfb\xda\x83\xe2\x16\x7c\xe5\x12\xf2\x85\x42\x47\x09\xbf\x1b\xa0\xbd\x15\x72\x6f\xbf\x72\x29\x26\x93\xef\xa1\x1f\xbf\x2b\x28\x2a\x64\x16\x77\x92\x9c\x3d\xb1\x2f\xba\x18\x4f\xe1\x5c\xd8\x9d\xd5\xde\x74\x99\x84\xc9\x49\xbe\xe8\x62\x44\x69\x3d\x72\x63\xea\xf4\xd0\xea\xef\x0e\xd7\xf1\x40\x17\x7e\x49\x8f\x7c\xd3\x7d\x10\xc2\x22\x51\x20\x2b\x95\x43\xab\xa0\xca\xa5\x99\x4c\x1a\xe7\x5e\x93\xfb\xb9\x47\x14\x19\x7b\x60\x4e\x0b\xcd\x2c\x1e\x74\xdd\x4e\x4b\xff\x62\x36\x4f\xa7\x7c\xca\xa7\xb3\xce\xea\x09\x6d\xc6\x8c\x26\xb7\xb3\x48\x9d\xd1\x00\xd1\x51\x5b\x11\x02\xf4\xae\xbc\x37\x76\x61\x18\x13\xe0\xa0\xb8\x14\xa1\x35\x82\x00\xe3\x7e\x80\x0c\x82\xcd\xae\x22\x5e\x55\xf3\x3a\x62\xef\xfa\x31\xa2\xd1\xd6\x85\x09\x4f\xd3\x34\xea\xa1\x7e\xd2\x9c\xff\x02\x36\x84\x8a\x1e\xd8\x27\x21\xd8\xe7\x3f\xff\xf8\x9f\xa6\x86\x7d\x28\x8d\xcc\x03\xd6\x65\x7c\xe0\x48\x79\xbf\xa3\xaf\xa6\xe7\xbc\xbd\xdf\x39\x45\x42\x5a\x2c\xc3\xe2\xee\x62\x0f\x64\xd3\xc7\x1e\x8f\x70\xd6\x5c\x2d\x05\xda\xc4\xe1\xc1\x54\xe0\x70\x80\xf1\xb1\xa7\xf4\xce\x34\xe1\x38\x14\xf4\x07\x38\xb6\x13\x02\x65\x89\x44\xf9\x2b\x9e\x72\xd9\x34\xee\xca\xd0\x35\x8d\xb0\xb4\xe8\xf2\x8b\xab\x59\x04\xb7\xc6\x5e\x07\xed\x62\x69\x10\xda\x16\x3c\x3e\xde\xda\x6e\xae\x12\x96\xde\x86\x05\xd9\x12\x18\xbe\xb8\x71\x75\x0f\xc3\x74\x85\x7f\x6d\x3b\x67\x5a\xdd\x5d\xa7\x37\x35\x26\xda\xe7\xce\x2b\x85\xb7\xbf\xb8\x71\x69\x77\x43\x3b\x8f\x2e\x3a\xaa\x4b\x30\xad\xcf\xca\x1a\x1c\xf6\x55\x19\x1c\x7f\x25\xa9\xd1\x46\x7d\x0c\x8d\x1a\x87\xa9\xc1\xca\xb0\x7f\xff\x65\x77\x9e\x55\xd8\xf8\x5a\xac\xb3\xe9\x9f\x00\x00\x00\xff\xff\x19\x93\xb7\xf9\x06\x09\x00\x00")

func vendorGithubComCloudfoundryBoshDeploymentAwsCpiYmlBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComCloudfoundryBoshDeploymentAwsCpiYml,
		"vendor/github.com/cloudfoundry/bosh-deployment/aws/cpi.yml",
	)
}

func vendorGithubComCloudfoundryBoshDeploymentAwsCpiYml() (*asset, error) {
	bytes, err := vendorGithubComCloudfoundryBoshDeploymentAwsCpiYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/cloudfoundry/bosh-deployment/aws/cpi.yml", size: 2310, mode: os.FileMode(420), modTime: time.Unix(1484943925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"vendor/github.com/cloudfoundry/bosh-deployment/bosh.yml":                                      vendorGithubComCloudfoundryBoshDeploymentBoshYml,
	"vendor/github.com/cloudfoundry/bosh-deployment/external-ip-not-recommended.yml":               vendorGithubComCloudfoundryBoshDeploymentExternalIpNotRecommendedYml,
	"vendor/github.com/cloudfoundry/bosh-deployment/external-ip-with-registry-not-recommended.yml": vendorGithubComCloudfoundryBoshDeploymentExternalIpWithRegistryNotRecommendedYml,
	"vendor/github.com/cloudfoundry/bosh-deployment/gcp/cpi.yml":                                   vendorGithubComCloudfoundryBoshDeploymentGcpCpiYml,
	"vendor/github.com/cloudfoundry/bosh-deployment/aws/cpi.yml":                                   vendorGithubComCloudfoundryBoshDeploymentAwsCpiYml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"vendor": &bintree{nil, map[string]*bintree{
		"github.com": &bintree{nil, map[string]*bintree{
			"cloudfoundry": &bintree{nil, map[string]*bintree{
				"bosh-deployment": &bintree{nil, map[string]*bintree{
					"aws": &bintree{nil, map[string]*bintree{
						"cpi.yml": &bintree{vendorGithubComCloudfoundryBoshDeploymentAwsCpiYml, map[string]*bintree{}},
					}},
					"bosh.yml":                                      &bintree{vendorGithubComCloudfoundryBoshDeploymentBoshYml, map[string]*bintree{}},
					"external-ip-not-recommended.yml":               &bintree{vendorGithubComCloudfoundryBoshDeploymentExternalIpNotRecommendedYml, map[string]*bintree{}},
					"external-ip-with-registry-not-recommended.yml": &bintree{vendorGithubComCloudfoundryBoshDeploymentExternalIpWithRegistryNotRecommendedYml, map[string]*bintree{}},
					"gcp": &bintree{nil, map[string]*bintree{
						"cpi.yml": &bintree{vendorGithubComCloudfoundryBoshDeploymentGcpCpiYml, map[string]*bintree{}},
					}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
