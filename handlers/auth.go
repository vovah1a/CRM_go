package handlers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt"
	"github.com/vovah1a/CRM_go/database"
	"github.com/vovah1a/CRM_go/models"
	password "github.com/vzglad-smerti/password_hash"
)

func GenerateJWT(user models.User) (string, error) {
	claims := golangJwt.MapClaims{
		"id":    user.ID,
		"email": user.Name,
		"exp":   time.Now().Add(time.Hour * 10).Unix(),
	}
	token := golangJwt.NewWithClaims(golangJwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_secret_key")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Auth(c *fiber.Ctx) error {
	user := models.User{Name: c.FormValue("Name")}

	if result := database.DB.Db.Where("name = ?", user.Name).Take(&user); result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": "логин и пароль не совпадают",
		})
	}

	hash_veriry, err := password.Verify(user.Password, c.FormValue("Password"))
	if err != nil || !hash_veriry {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": "логин и пароль не совпадают",
		})
	}

	tocken, err := GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"token": tocken})
}
