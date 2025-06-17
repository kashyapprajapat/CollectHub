package routes

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/kashyapprajapat/collecthub_api/controllers"

)

func SetupRoutes(app *fiber.App, db *mongo.Database) {
    controllers.InitUserController(db)
    controllers.InitBookController(db)


    //Home route
    app.Get("/", func(c *fiber.Ctx) error {
    htmlContent := `
        <html>
            <head>
                <title>CollectHub API</title>
            </head>
            <body>
                <h1>🚀 Welcome to the CollectHub API</h1>
                <p>This is the backend service for managing users and book collections.</p>
            </body>
        </html>
    `
     return c.Type("html").SendString(htmlContent)
    })

    api := app.Group("/api")



    // User Routes
    api.Post("/users", controllers.CreateUser)
    api.Get("/users", controllers.GetUsers)

    // Book Routes
    api.Post("/books", controllers.CreateBook)
    api.Get("/books/user/:userId", controllers.GetBooksByUser)
}
