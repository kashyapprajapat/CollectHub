package controllers

import (
    "context"
    "github.com/kashyapprajapat/collecthub_api/models"
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserController(db *mongo.Database) {
    userCollection = db.Collection("users")
}

func CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    res, err := userCollection.InsertOne(ctx, user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to insert user"})
    }

    return c.JSON(res)
}

func GetUsers(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := userCollection.Find(ctx, bson.M{})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to fetch users"})
    }
    var users []models.User
    if err = cursor.All(ctx, &users); err != nil {
        log.Fatal(err)
    }

    return c.JSON(users)
}
