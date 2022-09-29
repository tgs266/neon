package git

import (
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/utils"
)

func cloneRepo(c *gin.Context, repo string, credentials entities.Credentials) {

	dir := os.Getenv("NEON_HOME")

	repoPath := path.Base("repo")

	var auth transport.AuthMethod
	if credentials.UsingBasic() {
		auth = &http.BasicAuth{
			Username: credentials.Username,
			Password: utils.DecryptAES(c, utils.ReadKey(), credentials.Password),
		}
	} else {
		auth = &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: utils.DecryptAES(c, utils.ReadKey(), credentials.Token),
		}
	}
	options := &git.CloneOptions{
		URL:  repo,
		Auth: auth,
	}

	_, err := git.PlainClone(repoPath, false, options)
	fmt.Println(err)
}

func FillRepository(c *gin.Context, req api.CreateAppRequest) {
	credentials, err := store.CredentialsRepository().GetByName(req.CredentialName)
	if err != nil {
		errors.NewNotFound("credentials not found", err).Abort(c)
		return
	}

	cloneRepo(c, req.Repository, credentials)

}
