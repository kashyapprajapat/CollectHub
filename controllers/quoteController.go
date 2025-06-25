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

func GetQuoteByID(c *fiber.Ctx) error {
	quoteID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(quoteID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid quote ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var quote models.Quote
	filter := bson.M{"_id": objID}
	err = quoteCollection.FindOne(ctx, filter).Decode(&quote)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "quote not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch quote"})
	}

	return c.JSON(quote)
}

func UpdateQuote(c *fiber.Ctx) error {
	quoteID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(quoteID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid quote ID"})
	}

	var updateData models.Quote
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create update document, only include non-empty fields
	update := bson.M{}
	if updateData.Quote != "" {
		update["quote"] = updateData.Quote
	}
	if updateData.Author != "" {
		update["author"] = updateData.Author
	}
	if !updateData.UserID.IsZero() {
		update["user_id"] = updateData.UserID
	}

	if len(update) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "no fields to update"})
	}

	filter := bson.M{"_id": objID}
	updateDoc := bson.M{"$set": update}

	result, err := quoteCollection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update quote"})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "quote not found"})
	}

	return c.JSON(fiber.Map{
		"message":        "quote updated successfully",
		"modified_count": result.ModifiedCount,
	})
}

func DeleteQuote(c *fiber.Ctx) error {
	quoteID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(quoteID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid quote ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}
	result, err := quoteCollection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete quote"})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "quote not found"})
	}

	return c.JSON(fiber.Map{
		"message":       "quote deleted successfully",
		"deleted_count": result.DeletedCount,
	})
}