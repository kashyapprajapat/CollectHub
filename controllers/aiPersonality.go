package controllers

import (
    "context"
    "errors"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

// Top 3 📚 by id 
func GetTopBooksByUserID(db *mongo.Database, userID string) ([]fiber.Map, error) {
    // Parse string userID to ObjectID
    oid, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return nil, errors.New("invalid user ID format")
    }

    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := db.Collection("books")

    // Find books by user ID
    cursor, err := collection.Find(ctx, bson.M{"user_id": oid})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []fiber.Map
    count := 0

    for cursor.Next(ctx) {
        var book struct {
            BookName string `bson:"book_name"`
            Reason   string `bson:"reason"`
        }

        if err := cursor.Decode(&book); err != nil {
            continue // skip malformed documents
        }

        results = append(results, fiber.Map{
            "book_name": book.BookName,
            "reason":    book.Reason,
        })

        count++
        if count == 3 {
            break
        }
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

// Top 3 🎬 by id
func getTopMoviesByUserID(db *mongo.Database, userID string) ([]fiber.Map, error) {
    oid, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return nil, errors.New("invalid user ID format")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := db.Collection("movies")

    cursor, err := collection.Find(ctx, bson.M{"user_id": oid})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []fiber.Map
    count := 0

    for cursor.Next(ctx) {
        var movie struct {
            Title  string `bson:"title"`
            Type   string `bson:"type"`
            Reason string `bson:"reason"`
        }

        if err := cursor.Decode(&movie); err != nil {
            continue
        }

        results = append(results, fiber.Map{
            "title":  movie.Title,
            "type":   movie.Type,
            "reason": movie.Reason,
        })

        count++
        if count == 3 {
            break
        }
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

// Top 3 🐶 by id
func getTopPetByUserID(db *mongo.Database, userID string) ([]fiber.Map, error) {
    oid, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return nil, errors.New("invalid user ID format")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := db.Collection("pets")

    cursor, err := collection.Find(ctx, bson.M{"user_id": oid})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []fiber.Map
    count := 0

    for cursor.Next(ctx) {
        var pet struct {
            Name   string `bson:"name"`
            Reason string `bson:"reason"`
        }

        if err := cursor.Decode(&pet); err != nil {
            continue
        }

        results = append(results, fiber.Map{
            "name":   pet.Name,
            "reason": pet.Reason,
        })

        count++
        if count == 3 {
            break
        }
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

// Top 3 💬 by id
func getTopQuotesByUserID(db *mongo.Database, userID string) ([]fiber.Map, error) {
    oid, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return nil, errors.New("invalid user ID format")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := db.Collection("quotes")

    cursor, err := collection.Find(ctx, bson.M{"user_id": oid})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []fiber.Map
    count := 0

    for cursor.Next(ctx) {
        var quote struct {
            Quote string `bson:"quote"`
        }

        if err := cursor.Decode(&quote); err != nil {
            continue
        }

        results = append(results, fiber.Map{
            "quote": quote.Quote,
        })

        count++
        if count == 3 {
            break
        }
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

// Top 3 🍜🍕 by id
func getRecipeByUserID(db *mongo.Database, userID string) ([]fiber.Map, error) {
    oid, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return nil, errors.New("invalid user ID format")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := db.Collection("recipes")

    cursor, err := collection.Find(ctx, bson.M{"user_id": oid})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []fiber.Map
    count := 0

    for cursor.Next(ctx) {
        var recipe struct {
            Name   string `bson:"name"`
            Reason string `bson:"reason"`
        }

        if err := cursor.Decode(&recipe); err != nil {
            continue
        }

        results = append(results, fiber.Map{
            "name":   recipe.Name,
            "reason": recipe.Reason,
        })

        count++
        if count == 3 {
            break
        }
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

// Top 3 ✈️🧳🚢 by id
func getTravelByUserID(db *mongo.Database, userID string) ([]fiber.Map, error) {
    oid, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return nil, errors.New("invalid user ID format")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := db.Collection("travelbuddies")

    cursor, err := collection.Find(ctx, bson.M{"user_id": oid})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []fiber.Map
    count := 0

    for cursor.Next(ctx) {
        var travel struct {
            PlaceName string `bson:"place_name"`
            Reason    string `bson:"reason"`
        }

        if err := cursor.Decode(&travel); err != nil {
            continue
        }

        results = append(results, fiber.Map{
            "place_name": travel.PlaceName,
            "reason":     travel.Reason,
        })

        count++
        if count == 3 {
            break
        }
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}


