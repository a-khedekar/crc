package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/code-ready/crc/pkg/crc/machine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlainStatus(t *testing.T) {
	cacheDir, err := ioutil.TempDir("", "cache")
	require.NoError(t, err)
	defer os.RemoveAll(cacheDir)

	require.NoError(t, ioutil.WriteFile(filepath.Join(cacheDir, "crc.qcow2"), make([]byte, 10000), 0600))

	out := new(bytes.Buffer)
	assert.NoError(t, runStatus(out, &mockClient{}, cacheDir, ""))

	expected := `CRC VM:          Running
OpenShift:       Running (v4.5.1)
Disk Usage:      10GB of 20GB (Inside the CRC VM)
Cache Usage:     10kB
Cache Directory: %s
`
	assert.Equal(t, fmt.Sprintf(expected, cacheDir), out.String())
}

func TestJsonStatus(t *testing.T) {
	cacheDir, err := ioutil.TempDir("", "cache")
	require.NoError(t, err)
	defer os.RemoveAll(cacheDir)

	require.NoError(t, ioutil.WriteFile(filepath.Join(cacheDir, "crc.qcow2"), make([]byte, 10000), 0600))

	out := new(bytes.Buffer)
	assert.NoError(t, runStatus(out, &mockClient{}, cacheDir, jsonFormat))

	expected := `{
  "crcStatus": "Running",
  "openshiftStatus": "Running",
  "openshiftVersion": "4.5.1",
  "diskUsage": 10000000000,
  "diskSize": 20000000000,
  "cacheUsage": 10000,
  "cacheDir": "%s"
}
`
	assert.Equal(t, fmt.Sprintf(expected, cacheDir), out.String())
}

type mockClient struct{}

func (mockClient) Status(statusConfig machine.ClusterStatusConfig) (machine.ClusterStatusResult, error) {
	return machine.ClusterStatusResult{
		Name:             "crc",
		CrcStatus:        "Running",
		OpenshiftStatus:  "Running",
		OpenshiftVersion: "4.5.1",
		DiskUse:          10_000_000_000,
		DiskSize:         20_000_000_000,
		Success:          true,
	}, nil
}

func (mockClient) Exists(name string) (bool, error) {
	return true, nil
}