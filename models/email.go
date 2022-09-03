package models

import (
	"github.com/gofiber/fiber/v2"
	"regexp"
)

type EmailRequest struct {
	UserEmail string `json:"email"`
}

// ValidateFormat is a method to validate email using regex
func (e *EmailRequest) ValidateFormat() error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(e.UserEmail) {
		return fiber.ErrConflict
	}

	return nil
}

//func (e *EmailRequest) ValidateDns() error {}
