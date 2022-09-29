package api

type ApplyAppRequest struct {
	Name           string   `json:"name"`
	ReleaseChannel string   `json:"releaseChannel"`
	Products       []string `json:"products"`
}

type CreateAppRequest struct {
	Name           string   `json:"name"`
	ReleaseChannel string   `json:"releaseChannel"`
	Products       []string `json:"products"`
	Repository     string   `json:"repository"`
	CredentialName string   `json:"credentialName"`
}
