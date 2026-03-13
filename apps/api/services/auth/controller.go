package auth

import "github.com/gofiber/fiber/v3"

func AuthControllerHook(app *fiber.App) {
	authService := NewAuthService()

	app.Get("/auth/login", GetAuthHandler(authService))
}
