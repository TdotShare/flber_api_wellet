package main

import (
	"main/configs"
	"main/database"
	"main/routes"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @test Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @BasePath /
func main() {

	app := fiber.New()

	//config for customization
	app.Use(cors.New(configs.ConfigDefault))

	configs.InitDatabase()

	routes.AccountRoutes(app, database.DB)

	app.Get("/swagger/*", swagger.Handler) // default

	app.Listen(":3000")

}
