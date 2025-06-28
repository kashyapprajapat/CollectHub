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


func GetTravelByID(c *fiber.Ctx) error {
	travelID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(travelID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid travel ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var travel models.TravelBuddy
	err = travelCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&travel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "travel entry not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch travel entry"})
	}

	return c.JSON(travel)
}


func UpdateTravel(c *fiber.Ctx) error {
	travelID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(travelID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid travel ID"})
	}

	var updateData models.TravelBuddy
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create update document - only update non-empty fields
	update := bson.M{}
	if updateData.PlaceName != "" {
		update["place_name"] = updateData.PlaceName
	}
	if !updateData.DateVisited.IsZero() {
		update["date_visited"] = updateData.DateVisited
	}
	if updateData.Reason != "" {
		update["reason"] = updateData.Reason
	}
	if !updateData.UserID.IsZero() {
		update["user_id"] = updateData.UserID
	}

	if len(update) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "no fields to update"})
	}

	result, err := travelCollection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": update},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update travel entry"})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "travel entry not found"})
	}

	return c.JSON(fiber.Map{
		"message":        "travel entry updated successfully",
		"modified_count": result.ModifiedCount,
	})
}


func DeleteTravel(c *fiber.Ctx) error {
	travelID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(travelID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid travel ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := travelCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete travel entry"})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "travel entry not found"})
	}

	return c.JSON(fiber.Map{
		"message":       "travel entry deleted successfully",
		"deleted_count": result.DeletedCount,
	})
}