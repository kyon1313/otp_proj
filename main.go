package main

import (
	"log"
	"sample/database"
	"sample/model"
	"sample/route"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/add", route.AddUser)
	app.Post("/generateOtp", route.GenerateOtpEndpoint)
	app.Post("/validate", route.ValidateOtpEndpoint)
}

func main() {

	database.Migration()
	app := fiber.New()
	Routes(app)
	database.DB.AutoMigrate(&model.User{}, &model.OtpTable{})
	log.Fatal(app.Listen(":3000"))

}
