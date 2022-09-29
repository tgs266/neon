package api

type InstallConfig struct {
	Data string `json:"data"`
}

type InstallConfigCommit struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}
