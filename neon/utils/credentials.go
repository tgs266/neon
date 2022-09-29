package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"
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
