package api

import "github.com/tgs266/neon/neon/store/entities"

type CreateReleaseRequest struct {
	ProductName    string                `json:"productName"`
	ProductVersion string                `json:"productVersion"`
	ReleaseChannel string                `json:"releaseChannel"`
	HelmChart      string                `json:"helmChart"`
	Dependencies   []entities.Dependency `json:"dependencies"`
}
