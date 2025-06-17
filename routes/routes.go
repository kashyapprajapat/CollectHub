package routes

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/kashyapprajapat/collecthub_api/controllers"

)

func SetupRoutes(app *fiber.App, db *mongo.Database) {
    controllers.InitUserController(db)
    controllers.InitBookController(db)

    api := app.Group("/api")

    // User Routes
    api.Post("/users", controllers.CreateUser)
    api.Get("/users", controllers.GetUsers)

    // Book Routes
    api.Post("/books", controllers.CreateBook)
    api.Get("/books/user/:userId", controllers.GetBooksByUser)
}
