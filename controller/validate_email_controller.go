package controller

import (
	"github.com/XxThunderBlastxX/email-validator/models"
	"github.com/XxThunderBlastxX/email-validator/utils"
	"github.com/gofiber/fiber/v2"
)

// ValidateEmail is a handler to verify the email
func ValidateEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var email models.EmailRequest
		var resSuccess models.ResponseSuccess

		if err := ctx.BodyParser(&email); err != nil {
			parseErr := models.ResponseErr{Error: err.Error()}
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(parseErr.PresentErr())
		}

		// Checking the format of the email using regex
		if formatErr := email.ValidateFormat(); formatErr != nil {
			formatErrRes := models.ResponseErr{Error: formatErr.Error()}
			return ctx.Status(fiber.StatusConflict).JSON(formatErrRes.PresentErr())
		}

		domain := email.SplitEmail()

		// Checks for the mx-records for the domain
		if mxRecord := utils.MxRecord(domain); mxRecord == nil {
			resSuccess.MxRecord = nil
		} else {
			resSuccess.MxRecord = append(resSuccess.MxRecord, mxRecord...)
		}

		return ctx.Status(fiber.StatusOK).JSON(resSuccess.PresentSuccess())
	}
}
