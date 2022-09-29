package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/errors"
)

func GenerateNewKey() string {
	key := make([]byte, 32)

	rand.Read(key)
	return base64.StdEncoding.EncodeToString(key)
}

func WriteKey() error {
	dir := os.Getenv("NEON_HOME")
	if dir == "" {
		panic("must set NEON_HOME env")
	}
	key := GenerateNewKey()
	d1 := []byte(key)
	err := os.WriteFile(path.Join(dir, "keys"), d1, 0644)
	return err
}

func ReadKey() []byte {
	dir := os.Getenv("NEON_HOME")
	if dir == "" {
		panic("must set NEON_HOME env")
	}
	data, err := ioutil.ReadFile(path.Join(dir, "keys"))
	if err != nil {
		panic("couldnt read keyfile: " + err.Error())
	}
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		panic("couldnt read keyfile: " + err.Error())
	}
	return decoded
}

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
