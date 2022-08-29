package api

type CreateAppRequest struct {
	Name     string   `json:"name"`
	Products []string `json:"products"`
}
