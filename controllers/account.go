package controllers

import (
	"main/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AccountController struct {
	Db *gorm.DB
}

func (ac AccountController) ActionAll(c *fiber.Ctx) error {
	accountRepo := models.AccountsRepository{DB: ac.Db}
	account := accountRepo.GetAll()

	//var result []models.Account
	//ac.Db.Raw("SELECT * FROM wt_account").Scan(&result)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    account,
		"success": true,
	})
}

func (ac AccountController) ActionView(c *fiber.Ctx) error {

	id := c.Params("id")

	accountRepo := models.AccountsRepository{DB: ac.Db}
	account, err := accountRepo.GetOne(id)

	if err != nil {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    nil,
			"success": false,
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    account,
		"success": true,
	})
}

func (ac AccountController) ActionCreate(c *fiber.Ctx) error {
	accountRepo := models.AccountsRepository{DB: ac.Db}
	account := new(models.Account)

	if err := c.BodyParser(account); err != nil {
		return err
	}

	savedAccount := accountRepo.Save(*account)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    savedAccount,
		"success": true,
	})

}

func (ac AccountController) ActionDelete(c *fiber.Ctx) error {

	id := c.Params("id")
	accountRepo := models.AccountsRepository{DB: ac.Db}
	err := accountRepo.Delete(id)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "Delete Account Successfully",
		"success": true,
	})

}
