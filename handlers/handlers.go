package handlers

import (
	"boilerplate/database"
	"boilerplate/models"
	"fmt"
	"time"

	p "boilerplate/utils/password"
	t "boilerplate/utils/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type UserResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// Gets current user
func UserGet(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized",
		})
	}

	user := database.FindByToken(token)
	if user == nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}

	response := UserResponse{Email: user.Email, Token: user.Token}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"user":    response,
	})
}

// Authenticates a user
func UserLogin(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if (email == "") || (password == "") {
		return c.Status(422).JSON(fiber.Map{
			"success": false,
			"error":   "Email and password are required",
		})
	}

	user := database.FindByEmail(email)
	if user == nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}

	if p.Valid(password, user.Hash) {
		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = user.Token
		cookie.Expires = time.Now().Add(24 * time.Hour) // 1d ttl
		c.Cookie(cookie)

		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"token":   user.Token,
		})
	}
	return c.Status(401).JSON(fiber.Map{
		"success": false,
		"error":   "Invalid passphrase",
	})
}

func UserLogout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-24 * time.Hour) // -1d ttl
	c.Cookie(cookie)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Logged out",
	})
}

// Creates a new user
func UserCreate(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if (email == "") || (password == "") {
		return c.Status(422).JSON(fiber.Map{
			"success": false,
			"error":   "Email and password are required",
		})
	}

	if database.Exists(email) {
		return c.Status(409).JSON(fiber.Map{
			"success": false,
			"error":   "Email already exists",
		})
	}

	hash, err := p.New(password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Something went wrong. Please try again.",
		})
	}

	token := t.New(email, hash)
	user := &models.User{
		// Note: when writing to external database,
		// we can simply use - `Email: email`
		Email: utils.CopyString(email),
		Hash:  utils.CopyString(hash),
		Token: utils.CopyString(token),
	}
	database.Insert(user)

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = string(token)
	cookie.Expires = time.Now().Add(24 * time.Hour) // 1d ttl
	c.Cookie(cookie)

	fmt.Printf("User created: %s\nToken: %s\n", email, token)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"email":   user.Email,
		"token":   user.Token,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
