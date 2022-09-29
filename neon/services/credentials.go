package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/utils"
)

func EncryptAES(c *gin.Context, key []byte, plaintext string) string {
	block, err := aes.NewCipher(key)
	plainText := []byte(plaintext)
	if err != nil {
		errors.NewInternal("failed to encrypt", err).Abort(c)
		return ""
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		errors.NewInternal("failed to encrypt", err).Abort(c)
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return base64.RawStdEncoding.EncodeToString(cipherText)
}

func DecryptAES(c *gin.Context, key []byte, secure string) string {
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)

	if err != nil {
		errors.NewInternal("failed to decrypt", err).Abort(c)
		return ""
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		errors.NewInternal("failed to decrypt", err).Abort(c)
		return ""
	}

	if len(cipherText) < aes.BlockSize {
		errors.NewInternal("failed to decrypt", err).Abort(c)
		return ""
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}

func AddCredentials(c *gin.Context, req api.AddCredentialsRequest) api.AddCredentialsResponse {
	var e entities.Credentials
	if req.UsingBasic() {
		password := EncryptAES(c, utils.ReadKey(), req.Password)
		e = entities.Credentials{
			Name:     req.Name,
			Username: req.Username,
			Password: password,
		}
	} else {
		token := EncryptAES(c, utils.ReadKey(), req.Token)
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

func GetCredentials(c *gin.Context) []api.Credential {
	creds, err := store.CredentialsRepository().GetAll()
	if err != nil {
		errors.NewInternal("could not retrieve credentials", err).Abort(c)
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
