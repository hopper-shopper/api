package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/steschwa/hopper-analytics-api/controllers/compare"
	"github.com/steschwa/hopper-analytics-api/controllers/hoppers"
	"github.com/steschwa/hopper-analytics-api/controllers/prices"
	"github.com/steschwa/hopper-analytics-api/controllers/votes"
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

	server.Get("/hoppers", hoppers.NewRouteHandler(mongoClient))
	server.Get("/hoppers/compare", compare.NewRouteHandler(mongoClient))
	server.Get("/votes", votes.NewRouteHandler(mongoClient))
	server.Get("/prices", prices.NewRouteHandler(mongoClient))

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
