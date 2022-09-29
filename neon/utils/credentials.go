package utils

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func GenerateNewKey() string {
	key := make([]byte, 32)

	k, _ := rand.Read(key)
	return string(k)
}

func WriteKey() error {
	dir := os.Getenv("NEON_HOME")
	if dir == "" {
		panic("must set NEON_HOME env")
	}
	key := GenerateNewKey()
	fmt.Println(key)
	d1 := []byte(key)
	err := os.WriteFile(path.Join(dir, "keys"), d1, 0644)
	return err
}

func ReadKey() string {
	dir := os.Getenv("NEON_HOME")
	if dir == "" {
		panic("must set NEON_HOME env")
	}
	data, err := ioutil.ReadFile(path.Join(dir, "keys"))
	if err != nil {
		panic("couldnt read keyfile: " + err.Error())
	}
	return string(data)
}
