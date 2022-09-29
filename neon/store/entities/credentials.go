package entities

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/tgs266/neon/neon/utils"
	"github.com/uptrace/bun"
)

type Credentials struct {
	bun.BaseModel `bun:"table:credentials,alias:creds"`
	Name          string `bun:",pk" json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Token         string `json:"token"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
}

func (a Credentials) UsingBasic() bool {
	if a.Username != "" && a.Password != "" {
		return true
	}
	return false
}

func (a Credentials) GetGitCreds(c *gin.Context) transport.AuthMethod {
	var auth transport.AuthMethod
	if a.UsingBasic() {
		auth = &http.BasicAuth{
			Username: a.Username,
			Password: utils.DecryptAES(c, utils.ReadKey(), a.Password),
		}
	} else {
		auth = &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: utils.DecryptAES(c, utils.ReadKey(), a.Token),
		}
	}
	return auth
}
