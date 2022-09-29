package services

import (
	"bytes"
	"os/exec"

	"github.com/tgs266/neon/neon/store/entities"
)

func installUpdateHelmChart(namespace string, name string, release *entities.Release, pathToConfig string) (string, string, error) {
	cmd := exec.Command("helm", "upgrade", "-f", pathToConfig, name, release.HelmChart, "--version="+release.ProductVersion, "--namespace="+namespace, "-i", "--create-namespace")
	var out bytes.Buffer
	var out2 bytes.Buffer
	cmd.Stderr = &out
	cmd.Stdout = &out2

	err := cmd.Run()
	return out2.String(), out.String(), err
}

func deleteHelmRelease(namespace string, name string) (string, error) {
	cmd := exec.Command("helm", "uninstall", name, "--namespace="+namespace)
	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Run()
	return out.String(), err
}
