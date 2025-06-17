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

var recipeCollection *mongo.Collection

func InitRecipeController(db *mongo.Database) {
	recipeCollection = db.Collection("recipes")
}

func CreateRecipe(c *fiber.Ctx) error {
	var recipe models.Recipe
	if err := c.BodyParser(&recipe); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := recipeCollection.InsertOne(ctx, recipe)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to insert recipe"})
	}

	return c.JSON(res)
}

func GetRecipesByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": objID}
	cursor, err := recipeCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch recipes"})
	}

	var recipes []models.Recipe
	if err = cursor.All(ctx, &recipes); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error decoding recipes"})
	}

	return c.JSON(recipes)
}
