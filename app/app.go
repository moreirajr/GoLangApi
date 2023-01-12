package app

import (
	"dependencies/config"
	"dependencies/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type App struct {
	db     *gorm.DB
	config *config.Config
}

func Run(cfg *config.Config) {
	app := new(App)
	app.config = cfg

	app.db = db.Connect(cfg)
	db.MigrateDb(app.db)

	server := fiber.New()
	server.Use(limiter.New(limiter.Config{Max: 100}), cors.New(cors.ConfigDefault), logger.New(CreateLogger()))

	app.Routes(server)

	log.Fatal(server.Listen(app.config.Port))
}
