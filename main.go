package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/steschwa/hopper-analytics-api/controllers"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

const (
	PORT      = "PORT"
	MONGO_URI = "MONGO_URI"
)

func main() {
	mongoUri := os.Getenv(MONGO_URI)
	if mongoUri == "" {
		log.Fatalf("Missing environment variable %s\n", MONGO_URI)
	}

	mongoClient, err := db.Connect(mongoUri)
	if err != nil {
		log.Fatalln(err)
	}
	defer mongoClient.Disconnect(context.Background())

	server := fiber.New()
	server.Use(cors.New())

	server.Get("/hoppers", controllers.NewHoppersController(mongoClient))

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
