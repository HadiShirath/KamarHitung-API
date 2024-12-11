package main

import (
	"kamar-hitung/apps/auth"
	"kamar-hitung/apps/kecamatan"
	"kamar-hitung/apps/kelurahan"
	"kamar-hitung/apps/message"
	"kamar-hitung/apps/tps"
	"kamar-hitung/apps/user"
	"kamar-hitung/external/database"
	"kamar-hitung/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load configuration
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	// Connect to the database
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	if db != nil {
		log.Println("DB Connected")
	}

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		// Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	api := app.Group("/v1")

	// Configure CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Change this to the allowed origin
		// AllowOrigins:     "https://kamarhitung.id", // Change this to the allowed origin
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	app.Static("/images", "./public/images")

	// Initialize routes
	auth.Init(api, db)
	tps.Init(api, db)
	kelurahan.Init(api, db)
	kecamatan.Init(api, db)
	user.Init(api, db)
	message.Init(api, db)

	// Start server
	log.Fatal(app.Listen(config.Cfg.App.Port))
}
