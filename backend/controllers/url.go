package controllers

import (
	"context"
	"time"
	"urlshort/initializers"
	"urlshort/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func ShortenURL (c *fiber.Ctx) error {
	var data struct{
		LongURL string `json:"long_url"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error":"Data uncorrect"})
	}

	if data.LongURL == "" {
		return c.Status(400).JSON(fiber.Map{"error":"Url can't empty"})
	}

	shortID := uuid.New().String()[:6]
	url := models.URL{
		ID: shortID,
		LongURL: data.LongURL,
		CreateAt: time.Now(),
	}

	_, err := initializers.DB.InsertOne(context.TODO(), url)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error MongoDB"})

	}
	return c.JSON(fiber.Map{
		"short_url":"http://localhost:3000/"+shortID,
		"long_url": data.LongURL,
	})
}

// truy cap link da rut gon 
func RedirectURL (c *fiber.Ctx) error{
	shortID := c.Params("shortID")

	var result models.URL
	err := initializers.DB.FindOne(context.TODO(), bson.M{"_id":shortID}).Decode(&result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error":"Can't find url"})
	}
	return c.Redirect(result.LongURL, 301)
} 