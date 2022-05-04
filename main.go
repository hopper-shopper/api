package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/steschwa/hopper-analytics-api/controllers/hoppers"
	"github.com/steschwa/hopper-analytics-api/controllers/markets"
	"github.com/steschwa/hopper-analytics-api/controllers/prices"
	"github.com/steschwa/hopper-analytics-api/controllers/supply"
	"github.com/steschwa/hopper-analytics-api/controllers/transfers"
	"github.com/steschwa/hopper-analytics-api/controllers/user"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

const (
	PORT      = "PORT"
	MONGO_URI = "MONGO_URI"
	MONGO_DB  = "MONGO_DB"
)

func main() {
	mongoUri := os.Getenv(MONGO_URI)
	if mongoUri == "" {
		log.Fatalf("Missing environment variable %s\n", MONGO_URI)
	}

	databaseName := os.Getenv(MONGO_DB)
	if databaseName == "" {
		log.Fatalf("Missing environment variable %s\n", MONGO_DB)
	}

	initSentry()
	defer sentry.Flush(2 * time.Second)

	dbClient := db.NewMongoDbClient(databaseName)
	err := dbClient.Connect(mongoUri)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
	}
	defer dbClient.Disconnect()

	onChainClient, err := contracts.NewOnChainClient()
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
	}

	server := fiber.New()
	server.Use(cors.New())

	server.Get("/hoppers", hoppers.NewRouteHandler(dbClient))
	// server.Get("/votes", votes.NewRouteHandler(mongoClient))
	// server.Get("/base-shares", baseshares.NewRouteHandler(mongoClient))
	// server.Get("/base-shares/history", baseshares.NewHistoryRouteHandler(mongoClient))
	server.Post("/prices/historical", prices.NewHistoricalPriceRouteHandler(dbClient))
	server.Get("/prices", prices.NewLatestPriceRouteHandler(dbClient))
	server.Get("/market", markets.NewMarketHistoryByHopper(dbClient))
	server.Get("/transfers", transfers.NewHistoryRouteHandler())
	server.Get("/user/cap", user.NewUserCapRouteHandler(onChainClient, dbClient))
	server.Get("/user/earnings", user.NewUserEarnings(onChainClient, dbClient))
	server.Get("/supply", supply.NewRouteHandler(dbClient))

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

func initSentry() {
	if env := os.Getenv("ENV"); env == "production" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: "https://01da0a181e384b2f94fc5e93090bbeaf@o1202748.ingest.sentry.io/6328507",
		})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Using Sentry for error reporting")
	}
}
