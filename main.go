package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	baseshares "github.com/steschwa/hopper-analytics-api/controllers/base-shares"
	"github.com/steschwa/hopper-analytics-api/controllers/hoppers"
	"github.com/steschwa/hopper-analytics-api/controllers/markets"
	"github.com/steschwa/hopper-analytics-api/controllers/prices"
	"github.com/steschwa/hopper-analytics-api/controllers/transfers"
	"github.com/steschwa/hopper-analytics-api/controllers/user"
	"github.com/steschwa/hopper-analytics-api/controllers/votes"
	"github.com/steschwa/hopper-analytics-collector/contracts"
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

	onChainClient, err := contracts.NewOnChainClient()
	if err != nil {
		log.Fatalln(err)
	}

	server := fiber.New()
	server.Use(cors.New())

	server.Get("/hoppers", hoppers.NewRouteHandler(mongoClient))
	server.Get("/votes", votes.NewRouteHandler(mongoClient))
	server.Get("/base-shares", baseshares.NewRouteHandler(mongoClient))
	server.Get("/prices", prices.NewRouteHandler(mongoClient))
	server.Get("/market", markets.NewMarketHistoryRouteHandler(mongoClient))
	server.Get("/transfers", transfers.NewRouteHandler())
	server.Get("/user", user.NewRouteHandler(onChainClient))

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
