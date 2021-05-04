package routes

import (
	"main/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AccountRoutes(app *fiber.App, db *gorm.DB) {
	group := app.Group("/api/account")
	account := controllers.AccountController{Db: db}

	group.Get("/all", account.ActionAll)
	group.Get("/view/:id", account.ActionView)
	group.Post("/create", account.ActionCreate)
	group.Delete("/delete/:id", account.ActionDelete)

}
