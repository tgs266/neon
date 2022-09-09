package api

type ResourceList struct {
	Pods     []string `json:"pods"`
	Services []string `json:"services"`
}
