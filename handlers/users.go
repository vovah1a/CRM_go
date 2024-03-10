package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vovah1a/CRM_go/database"
	"github.com/vovah1a/CRM_go/models"
	password "github.com/vzglad-smerti/password_hash"
)

func Users(c *fiber.Ctx) error {
	user := []models.User{}

	database.DB.Db.Select("name", "ID").Find(&user)

	return c.Status(200).JSON(user)
}

func User(c *fiber.Ctx) error {
	user := models.User{ID: uint(c.QueryInt("id"))}

	if result := database.DB.Db.Select("name", "ID").Take(&user); result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).SendString("data not found")
	}

	return c.Status(200).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}
	hash, err := password.Hash(c.FormValue("Password"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}
	user.Password = hash

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(fiber.Map{"massage": "sacsess"})
}
