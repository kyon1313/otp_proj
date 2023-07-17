package route

import (
	"sample/database"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

func GenerateOtpEndpoint(c *fiber.Ctx) error {
	otp := new(model.OtpTable)
	if err := c.BodyParser(otp); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	if otp.SaveOtp() != "" {
		return c.JSON(&fiber.Map{
			"message": otp.SaveOtp(),
		})
	}
	return c.JSON(&fiber.Map{
		"message": "your otp is " + otp.Opt,
	})
}

func ValidateOtpEndpoint(c *fiber.Ctx) error {
	var (
		confirm_otp model.ValidateOtp
		otp         model.OtpTable
	)
	if err := c.BodyParser(&confirm_otp); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	database.DB.Debug().Raw("select * from otp_tables where user_id=?", confirm_otp.UserId).Find(&otp)
	if otp.Opt == "" {
		return c.JSON("User not exist")
	}
	success, message := otp.ValidateOtp(confirm_otp.ConfirmOtp)
	if !success {
		return c.JSON(&fiber.Map{
			"message": message,
		})
	}
	return c.JSON(&fiber.Map{
		"message": message,
	})

}
