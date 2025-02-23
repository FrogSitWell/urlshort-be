package main

import (
	"log"
	"urlshort/controllers"
	"urlshort/initializers"

	"github.com/gofiber/fiber/v2"
)

func main (){
	initializers.ConnectDB()

	app := fiber.New()
	app.Post("/shorten", controllers.ShortenURL)
	app.Get("/:shortID", controllers.RedirectURL)

	log.Println("server run at http://localhost:3000")
	app.Listen(":3000")
}