package main

import (
	"log"

	"github.com/tr1sm0s1n/fiber-postgres-api/controllers"
	"github.com/tr1sm0s1n/fiber-postgres-api/db"
	"github.com/tr1sm0s1n/fiber-postgres-api/middlewares"
	"github.com/tr1sm0s1n/fiber-postgres-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect the database")
	}

	db.AutoMigrate(&models.Certificate{})

	app := fiber.New()
	app.Use(logger.New())
	app.Use(func(ctx *fiber.Ctx) error {
		return middlewares.Authority(ctx)
	})

	app.Post("/create", func(ctx *fiber.Ctx) error {
		return controllers.CreateOne(ctx, db)
	})
	app.Get("/read", func(ctx *fiber.Ctx) error {
		return controllers.ReadAll(ctx, db)
	})
	app.Get("/read/:id", func(ctx *fiber.Ctx) error {
		return controllers.ReadOne(ctx, db)
	})
	app.Put("/update/:id", func(ctx *fiber.Ctx) error {
		return controllers.UpdateOne(ctx, db)
	})
	app.Delete("/delete/:id", func(ctx *fiber.Ctx) error {
		return controllers.DeleteOne(ctx, db)
	})
	app.Listen(":8080")
}
