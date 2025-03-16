package main

import (
	"log"
	"urlshort/controllers"
	"urlshort/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main (){
	initializers.ConnectDB()
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Cho phép tất cả các domain truy cập (có thể chỉnh sửa)
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type",
	}))
	app.Post("/shorten", controllers.ShortenURL)
	app.Get("/:shortID", controllers.RedirectURL)

	log.Println("server run at http://localhost:3000")
	app.Listen(":3000")
}