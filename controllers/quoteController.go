package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kashyapprajapat/collecthub_api/models"
)

var quoteCollection *mongo.Collection

func InitQuoteController(db *mongo.Database) {
	quoteCollection = db.Collection("quotes")
}

func CreateQuote(c *fiber.Ctx) error {
	var quote models.Quote
	if err := c.BodyParser(&quote); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := quoteCollection.InsertOne(ctx, quote)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to insert quote"})
	}

	return c.JSON(res)
}

func GetQuotesByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": objID}
	cursor, err := quoteCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch quotes"})
	}

	var quotes []models.Quote
	if err = cursor.All(ctx, &quotes); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error decoding quotes"})
	}

	return c.JSON(quotes)
}
