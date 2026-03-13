package main

import (
	"log"
	"rfmtransportes-api/services/auth"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo das variaveis de ambiente.")
	}
	app := fiber.New()

	auth.AuthControllerHook(app)

	log.Fatal(app.Listen(":3000"))
}
