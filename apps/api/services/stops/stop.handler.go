package stops

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func GetStopsHandler(stop IService) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req StopRequest
		c.Bind().JSON(&req)

		resp, err := stop.GetStops(req.InitialDate, req.FinalDate)
		if err != nil {
			c.SendStatus(400)
			fmt.Printf("Erro ao solicitar as paradas. %v", err.Error())
			return c.JSON(err.Error())
		}
		c.SendStatus(200)
		return c.JSON(resp)
	}
}
