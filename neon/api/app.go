package api

type ApplyAppRequest struct {
	Name           string   `json:"name"`
	ReleaseChannel string   `json:"releaseChannel"`
	Products       []string `json:"products"`
}
