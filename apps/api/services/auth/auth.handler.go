package auth

import "github.com/gofiber/fiber/v3"

func GetAuthHandler(service IService) fiber.Handler {
	return func(c fiber.Ctx) error {
		resp, err := service.Login()
		if err != nil {
			c.Status(400)
			return c.SendString(err.Error())
		}
		return c.JSON(resp)
	}
}
