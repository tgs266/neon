package services

import (
	"bytes"
	"os/exec"

	"github.com/tgs266/neon/neon/store/entities"
)

func installUpdateHelmChart(name string, release *entities.Release) (string, error) {
	cmd := exec.Command("helm", "upgrade", name, release.HelmChart, "--version="+release.ProductVersion, "-i")
	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Run()
	return out.String(), err
}
