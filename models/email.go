package models

import (
	"github.com/gofiber/fiber/v2"
	"regexp"
	"strings"
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

// SplitEmail split the email-id and returns domain
func (e *EmailRequest) SplitEmail() string {
	domain := strings.Split(e.UserEmail, "@")

	// domain[1] returns "google.com" from user@google.com i.e. domain
	return domain[1]
}

//func (e *EmailRequest) ValidateDns() error {}
