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

var movieCollection *mongo.Collection

func InitMovieController(db *mongo.Database) {
	movieCollection = db.Collection("movies")
}

func CreateMovie(c *fiber.Ctx) error {
	var movie models.Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := movieCollection.InsertOne(ctx, movie)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to insert movie"})
	}

	return c.JSON(res)
}

func GetMoviesByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": objID}
	cursor, err := movieCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch movies"})
	}

	var movies []models.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error decoding movies"})
	}

	return c.JSON(movies)
}
