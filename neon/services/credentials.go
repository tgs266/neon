package services

import (
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/utils"
)

func AddCredentials(c *gin.Context, req api.AddCredentialsRequest) api.AddCredentialsResponse {
	var e entities.Credentials
	if req.UsingBasic() {
		password := utils.EncryptAES(c, utils.ReadKey(), req.Password)
		e = entities.Credentials{
			Name:     req.Name,
			Username: req.Username,
			Password: password,
		}
	} else {
		token := utils.EncryptAES(c, utils.ReadKey(), req.Token)
		e = entities.Credentials{
			Name:  req.Name,
			Token: token,
		}
	}
	if err := store.CredentialsRepository().Insert(e); err != nil {
		errors.NewInternal("failed to store credentials", err).Panic()
		return api.AddCredentialsResponse{}
	}
	return api.AddCredentialsResponse{
		Name: e.Name,
	}
}

func GetCredentials(c *gin.Context) []api.Credential {
	creds, err := store.CredentialsRepository().GetAll()
	if err != nil {
		errors.NewInternal("could not retrieve credentials", err).Panic()
		return nil
	}

	outCreds := []api.Credential{}
	for _, c := range creds {
		outCreds = append(outCreds, api.Credential{
			Name:      c.Name,
			BasicAuth: c.UsingBasic(),
			TokenAuth: !c.UsingBasic(),
		})
	}
	return outCreds
}
