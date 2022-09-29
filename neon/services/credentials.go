package services

import (
	"crypto/aes"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/utils"
)

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(errors.NewInternal("failed to encrypt", err))
	}
	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func AddCredentials(c *gin.Context, req api.AddCredentialsRequest) api.AddCredentialsResponse {
	var e entities.Credentials
	if req.UsingBasic() {
		password := EncryptAES(utils.ReadKey(), req.Password)
		e = entities.Credentials{
			Name:     req.Name,
			Username: req.Username,
			Password: password,
		}
	} else {
		token := EncryptAES(utils.ReadKey(), req.Token)
		e = entities.Credentials{
			Name:  req.Name,
			Token: token,
		}
	}
	if err := store.CredentialsRepository().Insert(e); err != nil {
		errors.NewInternal("failed to store credentials", err).Abort(c)
		return api.AddCredentialsResponse{}
	}
	return api.AddCredentialsResponse{
		Name: e.Name,
	}
}
