package main

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/vovah1a/CRM_go/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Post("/auth", handlers.Auth)
	user := app.Group("/user")

	user.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_secret_key")),
		},
	}))
	user.Get("/get-all", handlers.Users)
	user.Get("", handlers.User)
	user.Post("", handlers.CreateUser)
}
