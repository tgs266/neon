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

func (a AddCredentialsRequest) Validate(c *gin.Context) {
	if a.Name == "" || strings.Contains(a.Name, " ") {
		errors.NewBadRequest("name cannot be empty or contain spaces", nil).Abort(c)
		return
	}

	if a.Username != "" && a.Password == "" || a.Username == "" && a.Password != "" {
		errors.NewBadRequest("if using username/password auth, must provide both", nil).Abort(c)
		return
	}

	if a.Username == "" && a.Password == "" && a.Token == "" {
		errors.NewBadRequest("must provide authentication parameters", nil).Abort(c)
		return
	}

}

func (a AddCredentialsRequest) UsingBasic() bool {
	if a.Username != "" && a.Password != "" {
		return true
	}
	return false
}
