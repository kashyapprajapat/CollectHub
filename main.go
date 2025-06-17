package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "github.com/kashyapprajapat/collecthub_api/routes"
    "github.com/gofiber/fiber/v2/middleware/cors" 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DB")
    port := os.Getenv("PORT")

    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // ✅ Log successful MongoDB connection
    fmt.Println("✅ MongoDB database connected successfully.")

    db := client.Database(dbName)

    app := fiber.New()
   
    // 🔓 Enable CORS for all origins
    app.Use(cors.New())
    
    routes.SetupRoutes(app, db)

    // ✅ Log server startup info
    fmt.Printf("🚀 CollectHub API running on port: %s\n", port)

    log.Fatal(app.Listen(":" + port))
}
