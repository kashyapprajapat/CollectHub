package controllers

import (
    "context"
    "github.com/kashyapprajapat/collecthub_api/models"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection

func InitBookController(db *mongo.Database) {
    bookCollection = db.Collection("books")
}

func CreateBook(c *fiber.Ctx) error {
    var book models.Book
    if err := c.BodyParser(&book); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    res, err := bookCollection.InsertOne(ctx, book)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to insert book"})
    }

    return c.JSON(res)
}

func GetBooksByUser(c *fiber.Ctx) error {
    userID := c.Params("userId")
    objID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid user ID"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"user_id": objID}
    cursor, err := bookCollection.Find(ctx, filter)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to fetch books"})
    }

    var books []models.Book
    if err = cursor.All(ctx, &books); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "error decoding books"})
    }

    return c.JSON(books)
}
