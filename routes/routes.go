package routes

import (
	"github.com/XxThunderBlastxX/email-validator/controller"
	"github.com/XxThunderBlastxX/email-validator/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Router(app *fiber.App) {
	//Initial Route
	app.Get("/", middleware.RateLimiting(), controller.InitialRoute())

	//Fiber Monitor
	app.Get("/monitor", middleware.RateLimiting(), monitor.New(monitor.Config{Title: "Chamting-API"}))

	// Email validator route
	app.Post("/valid", middleware.RateLimiting(), controller.ValidateEmail())

}
