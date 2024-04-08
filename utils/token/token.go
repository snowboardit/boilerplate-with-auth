package token

import (
	"boilerplate/utils/env"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func New(email, hash string) string {
	secret := env.GetSecret()
	data := fmt.Sprintf("%s%s%s", email, secret, hash)

	return generate(data)
}

func generate(data string) string {
	hash := md5.New()
	io.WriteString(hash, data)
	token := hex.EncodeToString(hash.Sum(nil))

	return token
}
