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
