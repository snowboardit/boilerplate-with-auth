package env

import "os"

const secret = "SECRET"

func GetSecret() string {
	s := os.Getenv(secret)

	return s
}

func Verify() {
	s := GetSecret()
	if s == "" {
		panic("Secret key not found")
	}
}
