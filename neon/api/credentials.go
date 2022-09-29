package api

import (
	"strings"

	"k8s.io/apimachinery/pkg/api/errors"
)

type AddCredentialsRequest struct {
	Name string `json:"name"`

	Username string `json:"username"`
	Password string `json:"password"`

	Token string `json:"token"`
}

func (a AddCredentialsRequest) Validate() {
	if a.Name == "" || strings.Contains(a.Name, " ") {
		panic(errors.NewBadRequest("name cannot be empty or contain spaces"))
	}

	if a.Username != "" && a.Password == "" || a.Username == "" && a.Password != "" {
		panic(errors.NewBadRequest("if using username/password auth, must provide both"))
	}

	if a.Username == "" && a.Password == "" && a.Token == "" {
		panic(errors.NewBadRequest("must provide authentication parameters"))
	}

}

func (a AddCredentialsRequest) UsingBasic() bool {
	if a.Username != "" && a.Password != "" {
		return true
	}
	return false
}
