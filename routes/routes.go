package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
    "runtime"
	"time"
	"fmt"
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
				<h1>Welcome to the CollectHub API</h1>
				<p>This is the backend service for managing users and collections (books, movies, quotes, pets, travel, etc.).</p>
			</body>
		</html>
		`
		return c.Type("html").SendString(htmlContent)
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	var startTime = time.Now()
	
	//System health route
    app.Get("/health", func(c *fiber.Ctx) error {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		uptime := time.Since(startTime).Truncate(time.Second)

		htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>System Health - CollectHub</title>
			<style>
				body {
					font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
					background-color: #f0f2f5;
					margin: 0;
					padding: 30px;
				}
				.container {
					background: white;
					max-width: 700px;
					margin: auto;
					padding: 30px;
					border-radius: 12px;
					box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
				}
				h1 {
					color: #2c3e50;
					margin-bottom: 20px;
				}
				p {
					font-size: 16px;
					color: #34495e;
					margin: 8px 0;
				}
				.label {
					font-weight: bold;
				}
				.footer {
					margin-top: 30px;
					font-size: 14px;
					color: #888;
					text-align: center;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>System Health Dashboard</h1>
				<p><span class="label">Go Version:</span> %s</p>
				<p><span class="label">Num CPU:</span> %d</p>
				<p><span class="label">Uptime:</span> %s</p>
				<p><span class="label">Goroutines:</span> %d</p>
				<p><span class="label">GC Cycles:</span> %d</p>
				<hr>
				<p><span class="label">Memory Allocated:</span> %s</p>
				<p><span class="label">Total Memory Allocated:</span> %s</p>
				<p><span class="label">System Memory Obtained:</span> %s</p>
			</div>
			<div class="footer">
				&copy; 2025 CollectHub. All rights reserved.
			</div>
		</body>
		</html>
		`, runtime.Version(), runtime.NumCPU(), uptime, runtime.NumGoroutine(), memStats.NumGC,
			formatBytes(memStats.Alloc), formatBytes(memStats.TotalAlloc), formatBytes(memStats.Sys))

		return c.Type("html").SendString(htmlContent)
	})

	api := app.Group("/api")

	// User Routes
    api.Post("/users", controllers.CreateUser)
    api.Get("/users", controllers.GetUsers)
    api.Post("/users/login", controllers.LoginUser)

	// Book Routes
	api.Post("/books", controllers.CreateBook)
	api.Get("/books/user/:userId", controllers.GetBooksByUser)
	api.Get("/books/:id", controllers.GetBookByID)
    api.Put("/books/:id", controllers.UpdateBook)
    api.Delete("/books/:id", controllers.DeleteBook)

	// Recipe Routes
	api.Post("/recipes", controllers.CreateRecipe)
	api.Get("/recipes/user/:userId", controllers.GetRecipesByUser)
	api.Get("/recipes/:id", controllers.GetRecipeByID)               
    api.Put("/recipes/:id", controllers.UpdateRecipe)        
	api.Delete("/recipes/:id", controllers.DeleteRecipe)         

	// Movie Routes
	api.Post("/movies", controllers.CreateMovie)
	api.Get("/movies/user/:userId", controllers.GetMoviesByUser)
	api.Get("/movies/:id", controllers.GetMovieByID)
    api.Put("/movies/:id", controllers.UpdateMovie)
    api.Delete("/movies/:id", controllers.DeleteMovie)

	// Quote Routes
	api.Post("/quotes", controllers.CreateQuote)
	api.Get("/quotes/user/:userId", controllers.GetQuotesByUser)
	api.Get("/quotes/:id", controllers.GetQuoteByID)
    api.Put("/quotes/:id", controllers.UpdateQuote)
    api.Delete("/quotes/:id", controllers.DeleteQuote)

	// Pet Routes
	api.Post("/pets", controllers.CreatePet)
	api.Get("/pets/user/:userId", controllers.GetPetsByUser)
	api.Get("/pets/:id", controllers.GetPetByID)
    api.Put("/pets/:id", controllers.UpdatePet)
    api.Delete("/pets/:id", controllers.DeletePet)

	// Travel Routes
	api.Post("/travels", controllers.CreateTravel)
	api.Get("/travels/user/:userId", controllers.GetTravelsByUser)
	api.Get("/travels/:id", controllers.GetTravelByID)             
    api.Put("/travels/:id", controllers.UpdateTravel)                
    api.Delete("/travels/:id", controllers.DeleteTravel)

	// 🤖 AI Personality Analysis Route
	api.Post("/aipersonality/analysis", controllers.GetAIPersonalityAnalysis(db))
}


// Helper functions
func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}