package stops

import "github.com/gofiber/fiber/v3"

func ControllerHook(app *fiber.App) {
	stopService := NewStopService()

	app.Post("/stops", GetStopsHandler(stopService))
}
