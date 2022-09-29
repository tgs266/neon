package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/errors"
)

type AddCredentialsRequest struct {
	Name string `json:"name"`

	Username string `json:"username"`
	Password string `json:"password"`

	Token string `json:"token"`
}

type AddCredentialsResponse struct {
	Name string `json:"name"`
}

type Credential struct {
	Name      string `json:"name"`
	BasicAuth bool   `json:"basicAuth"`
	TokenAuth bool   `json:"tokenAuth"`
}

func (a AddCredentialsRequest) Validate(c *gin.Context) {
	if a.Name == "" || strings.Contains(a.Name, " ") {
		panic(errors.NewBadRequest("name cannot be empty or contain spaces", nil))
		return
	}

	if a.Username != "" && a.Password == "" || a.Username == "" && a.Password != "" {
		panic(errors.NewBadRequest("if using username/password auth, must provide both", nil))
		return
	}

	if a.Username == "" && a.Password == "" && a.Token == "" {
		panic(errors.NewBadRequest("must provide authentication parameters", nil))
		return
	}

}

func (a AddCredentialsRequest) UsingBasic() bool {
	if a.Username != "" && a.Password != "" {
		return true
	}
	return false
}
