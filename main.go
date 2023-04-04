package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mr-emerald-wolf/mailer-go/initializers"
	"github.com/mr-emerald-wolf/mailer-go/utils"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	// Api rate limiter
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	fmt.Println("Server Started")

	app.Get("/ping", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(fiber.Map{
			"status":  "true",
			"message": "pong",
		})

	})

	app.Post("/send", func(c *fiber.Ctx) error {

		// Get request body and bind to payload
		var payload *utils.SendEmailRequest
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		// Validate struct
		errors := utils.ValidateStruct(payload)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}
		// Send email
		err := utils.SendMail(payload.Subject, payload.Body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "true", "message": "Email sent successfully"})
	})

	// Add Routes
	log.Fatal(app.Listen(":8081"))
}
