package auth

import "github.com/gofiber/fiber/v3"

func ControllerHook(app *fiber.App) {
	authService := NewAuthService()

	app.Get("/auth/login", GetAuthHandler(authService))
}
