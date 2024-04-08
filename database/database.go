package database

import (
	"boilerplate/models"
	"fmt"
	"sync"
)

var (
	db []*models.User
	mu sync.Mutex
)

// Connect with database
func Connect() {
	db = make([]*models.User, 0)
	fmt.Println("Connected with Database")
}

func Insert(user *models.User) {
	mu.Lock()
	db = append(db, user)
	mu.Unlock()
}

func Get() []*models.User {
	return db
}

func FindByEmail(email string) *models.User {
	for _, user := range db {
		if user.Email == email {
			return user
		}
	}
	return nil
}

func FindByToken(token string) *models.User {
	for _, user := range db {
		if user.Token == token {
			return user
		}
	}
	return nil
}

func Exists(email string) bool {
	for _, user := range db {
		if user.Email == email {
			return true
		}
	}
	return false
}
