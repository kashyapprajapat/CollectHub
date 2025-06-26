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

var petCollection *mongo.Collection

func InitPetController(db *mongo.Database) {
	petCollection = db.Collection("pets")
}

func CreatePet(c *fiber.Ctx) error {
	var pet models.Pet
	if err := c.BodyParser(&pet); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := petCollection.InsertOne(ctx, pet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to insert pet"})
	}

	return c.JSON(res)
}

func GetPetsByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": objID}
	cursor, err := petCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch pets"})
	}

	var pets []models.Pet
	if err = cursor.All(ctx, &pets); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error decoding pets"})
	}

	return c.JSON(pets)
}

func GetPetByID(c *fiber.Ctx) error {
	petID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid pet ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var pet models.Pet
	filter := bson.M{"_id": objID}
	err = petCollection.FindOne(ctx, filter).Decode(&pet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "pet not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch pet"})
	}

	return c.JSON(pet)
}

func UpdatePet(c *fiber.Ctx) error {
	petID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid pet ID"})
	}

	var updateData models.Pet
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create update document, only include non-empty fields
	update := bson.M{}
	if updateData.Name != "" {
		update["name"] = updateData.Name
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

	filter := bson.M{"_id": objID}
	updateDoc := bson.M{"$set": update}

	result, err := petCollection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update pet"})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "pet not found"})
	}

	return c.JSON(fiber.Map{
		"message":        "pet updated successfully",
		"modified_count": result.ModifiedCount,
	})
}

func DeletePet(c *fiber.Ctx) error {
	petID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid pet ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}
	result, err := petCollection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete pet"})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "pet not found"})
	}

	return c.JSON(fiber.Map{
		"message":       "pet deleted successfully",
		"deleted_count": result.DeletedCount,
	})
}