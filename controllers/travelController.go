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

var travelCollection *mongo.Collection

func InitTravelController(db *mongo.Database) {
	travelCollection = db.Collection("travels")
}

func CreateTravel(c *fiber.Ctx) error {
	var travel models.TravelBuddy
	if err := c.BodyParser(&travel); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := travelCollection.InsertOne(ctx, travel)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to insert travel entry"})
	}

	return c.JSON(res)
}

func GetTravelsByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": objID}
	cursor, err := travelCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch travel entries"})
	}

	var travels []models.TravelBuddy
	if err = cursor.All(ctx, &travels); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error decoding travel entries"})
	}

	return c.JSON(travels)
}
