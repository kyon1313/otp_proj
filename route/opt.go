package route

import (
	"fmt"
	"sample/database"
	"sample/helper"
	"sample/model"

	"github.com/gofiber/fiber/v2"
)

func GenerateOtpEndpoint(c *fiber.Ctx) error {
	otp := new(model.OtpTable)
	var message string
	if err := c.BodyParser(otp); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	if otp.SaveOtp() != "" {
		message = otp.SaveOtp()
		helper.LogOTPAction("Generate", otp.Opt, fmt.Sprintf("%d", otp.UserId), message)
		return c.JSON(&fiber.Map{
			"success": false,
			"message": message,
		})
	}

	message = "your otp is " + otp.Opt
	//the otp.userID will be replace with the mobile number

	helper.LogOTPAction("Generate", otp.Opt, fmt.Sprintf("%d", otp.UserId), message)
	return c.JSON(&fiber.Map{
		"success": true,
		"message": message,
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
		return c.JSON(&fiber.Map{
			"success": false,
			"message": "user not exist",
		})
	}
	success, message := otp.ValidateOtp(confirm_otp.ConfirmOtp)
	if !success {
		helper.LogOTPAction("Validate", confirm_otp.ConfirmOtp, fmt.Sprintf("%d", otp.UserId), message)
		return c.JSON(&fiber.Map{
			"success": false,
			"message": message,
		})
	}
	helper.LogOTPAction("Validate", confirm_otp.ConfirmOtp, fmt.Sprintf("%d", otp.UserId), message)
	return c.JSON(&fiber.Map{
		"success": true,
		"message": message,
	})

}
