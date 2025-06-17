package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kashyapprajapat/collecthub_api/controllers"
)

func SetupRoutes(app *fiber.App, db *mongo.Database) {
	// Initialize Controllers
	controllers.InitUserController(db)
	controllers.InitBookController(db)
	controllers.InitRecipeController(db)
	controllers.InitMovieController(db)
	controllers.InitQuoteController(db)
	controllers.InitPetController(db)
	controllers.InitTravelController(db)

	// Home Route
	app.Get("/", func(c *fiber.Ctx) error {
		htmlContent := `
		<html>
			<head>
				<title>CollectHub API</title>
			</head>
			<body>
				<h1>🚀 Welcome to the CollectHub API</h1>
				<p>This is the backend service for managing users and collections (books, movies, quotes, pets, travel, etc.).</p>
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

	// Recipe Routes
	api.Post("/recipes", controllers.CreateRecipe)
	api.Get("/recipes/user/:userId", controllers.GetRecipesByUser)

	// Movie Routes
	api.Post("/movies", controllers.CreateMovie)
	api.Get("/movies/user/:userId", controllers.GetMoviesByUser)

	// Quote Routes
	api.Post("/quotes", controllers.CreateQuote)
	api.Get("/quotes/user/:userId", controllers.GetQuotesByUser)

	// Pet Routes
	api.Post("/pets", controllers.CreatePet)
	api.Get("/pets/user/:userId", controllers.GetPetsByUser)

	// Travel Routes
	api.Post("/travels", controllers.CreateTravel)
	api.Get("/travels/user/:userId", controllers.GetTravelsByUser)
}
