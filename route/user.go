package route

import (
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

func AddUser(ctx *fiber.Ctx) error {
	user := new(model.User)
	ctx.BodyParser(user)
	err := user.HashPassword()
	if err != nil {
		return ctx.JSON(err)
	}
	message := user.ValidateUsername()
	return ctx.JSON(&fiber.Map{
		"message": message,
	})
}
