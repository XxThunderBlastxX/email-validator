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

		// Parsing the data from body to the route
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
			resSuccess.ContainsMx = false
		} else {
			resSuccess.MxRecord = append(resSuccess.MxRecord, mxRecord...)
			resSuccess.ContainsMx = true
		}

		// Check if it contains spf record
		containsSpf, spf := utils.SpfRecord(domain)
		if containsSpf {
			resSuccess.ContainsSpf = true
			resSuccess.Spf = spf
		} else {
			resSuccess.ContainsSpf = false
			resSuccess.Spf = spf
		}

		// Check if it contains dmarc record
		containsDmarc, dmarc := utils.DmarcRecord(domain)
		if containsDmarc {
			resSuccess.ContainsDmarc = true
			resSuccess.Dmarc = dmarc
		} else {
			resSuccess.ContainsDmarc = false
			resSuccess.Dmarc = dmarc
		}

		return ctx.Status(fiber.StatusOK).JSON(resSuccess.PresentSuccess())
	}
}
