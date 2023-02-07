package main

import (
	"log"

	"github.com/Dashaa011/fiber-tut/pkg/common/config"

	"github.com/Dashaa011/fiber-tut/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db.Init(c.DBUrl)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})

	// books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
