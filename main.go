package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/steschwa/hopper-analytics-api/controllers"
)

const (
	PORT = "PORT"
)

func main() {
	server := fiber.New()
	server.Use(cors.New())

	server.Get("/markets/:adventure", controllers.NewMarketsByAdventureController())

	server.Listen(getServerAddress())
}

func getServerAddress() string {
	port := os.Getenv(PORT)

	if port == "" {
		log.Printf("%s is not present in the environment, using Port 5001 instead", PORT)
		return ":5001"
	}

	return fmt.Sprintf(":%s", port)
}
