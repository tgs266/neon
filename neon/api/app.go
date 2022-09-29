package api

type ApplyAppRequest struct {
	Name           string   `json:"name"`
	ReleaseChannel string   `json:"releaseChannel"`
	Products       []string `json:"products"`
}

type CreateAppRequest struct {
	Name           string   `json:"name" yaml:"name"`
	Products       []string `json:"products" yaml:"products"`
	Repository     string   `json:"repository" yaml:"repository"`
	CredentialName string   `json:"credentialName" yaml:"credentialName"`
}
