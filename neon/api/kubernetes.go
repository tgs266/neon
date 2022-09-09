package api

type ResourceList struct {
	Pods     []string `json:"pods"`
	Services []string `json:"services"`
}

type PodStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
