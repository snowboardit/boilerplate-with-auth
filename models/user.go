package models

// User model
type User struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
	Token string `json:"token"`
}
