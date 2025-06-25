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

// Book collection init
func InitBookController(db *mongo.Database) {
    bookCollection = db.Collection("books")
}

// Create book
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

// Get books by user id
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

// GetBookByID gets a single book by ID
func GetBookByID(c *fiber.Ctx) error {
    bookID := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(bookID)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid book ID"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var book models.Book
    filter := bson.M{"_id": objID}
    err = bookCollection.FindOne(ctx, filter).Decode(&book)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(404).JSON(fiber.Map{"error": "book not found"})
        }
        return c.Status(500).JSON(fiber.Map{"error": "failed to fetch book"})
    }

    return c.JSON(book)
}

// UpdateBook updates a book by ID
func UpdateBook(c *fiber.Ctx) error {
    bookID := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(bookID)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid book ID"})
    }

    var updateData models.Book
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Create update document, only include non-empty fields
    update := bson.M{}
    if updateData.BookName != "" {
        update["book_name"] = updateData.BookName
    }
    if updateData.Author != "" {
        update["author"] = updateData.Author
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

    result, err := bookCollection.UpdateOne(ctx, filter, updateDoc)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to update book"})
    }

    if result.MatchedCount == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "book not found"})
    }

    return c.JSON(fiber.Map{
        "message": "book updated successfully",
        "modified_count": result.ModifiedCount,
    })
}

// DeleteBook deletes a book by ID
func DeleteBook(c *fiber.Ctx) error {
    bookID := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(bookID)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid book ID"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"_id": objID}
    result, err := bookCollection.DeleteOne(ctx, filter)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to delete book"})
    }

    if result.DeletedCount == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "book not found"})
    }

    return c.JSON(fiber.Map{
        "message": "book deleted successfully",
        "deleted_count": result.DeletedCount,
    })
}