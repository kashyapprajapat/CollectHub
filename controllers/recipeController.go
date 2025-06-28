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


func GetRecipeByID(c *fiber.Ctx) error {
	recipeID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid recipe ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var recipe models.Recipe
	err = recipeCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&recipe)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "recipe not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch recipe"})
	}

	return c.JSON(recipe)
}


func UpdateRecipe(c *fiber.Ctx) error {
	recipeID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid recipe ID"})
	}

	var updateData models.Recipe
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create update document - only update non-empty fields
	update := bson.M{}
	if updateData.Name != "" {
		update["name"] = updateData.Name
	}
	if updateData.Ingredients != "" {
		update["ingredients"] = updateData.Ingredients
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

	result, err := recipeCollection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": update},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update recipe"})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "recipe not found"})
	}

	return c.JSON(fiber.Map{
		"message":        "recipe updated successfully",
		"modified_count": result.ModifiedCount,
	})
}


func DeleteRecipe(c *fiber.Ctx) error {
	recipeID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid recipe ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := recipeCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete recipe"})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "recipe not found"})
	}

	return c.JSON(fiber.Map{
		"message":       "recipe deleted successfully",
		"deleted_count": result.DeletedCount,
	})
}