package services

import (
	"bytes"
	"os/exec"

	"github.com/tgs266/neon/neon/store/entities"
)

func installUpdateHelmChart(namespace string, name string, release *entities.Release) (string, error) {
	cmd := exec.Command("helm", "upgrade", name, release.HelmChart, "--version="+release.ProductVersion, "--namespace="+namespace, "-i", "--create-namespace")
	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Run()
	return out.String(), err
}
