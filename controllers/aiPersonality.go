package controllers

import (
    "context"
    "errors"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

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

    collection := db.Collection("travels")

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



// UserData represents all collected user data
type UserData struct {
	Books   []fiber.Map `json:"books"`
	Movies  []fiber.Map `json:"movies"`
	Pets    []fiber.Map `json:"pets"`
	Quotes  []fiber.Map `json:"quotes"`
	Recipes []fiber.Map `json:"recipes"`
	Travel  []fiber.Map `json:"travel"`
}

// GeminiRequest represents the request structure for Gemini API
type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

// GeminiResponse represents the response structure from Gemini API
type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// PersonalityAnalysisRequest represents the request body
type PersonalityAnalysisRequest struct {
	UserID string `json:"user_id"`
}

// GetAIPersonalityAnalysis is the main controller function
func GetAIPersonalityAnalysis(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body
		var req PersonalityAnalysisRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Validate user ID
		if req.UserID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "User ID is required",
			})
		}

		// Collect user data using goroutines
		userData, err := collectUserDataConcurrently(db, req.UserID)
		if err != nil {
			log.Printf("Error collecting user data: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to collect user data",
			})
		}

		// Generate personality analysis using Gemini API
		personalityAnalysis, err := generatePersonalityAnalysis(userData)
		if err != nil {
			log.Printf("Error generating personality analysis: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate personality analysis",
			})
		}

		// Return the result
		return c.JSON(fiber.Map{
			"success":             true,
			"user_id":            req.UserID,
			"personality_analysis": personalityAnalysis,
			"data_collected":      userData,
		})
	}
}

// collectUserDataConcurrently uses goroutines to fetch data from all collections
func collectUserDataConcurrently(db *mongo.Database, userID string) (*UserData, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	userData := &UserData{}
	errors := make([]error, 0)

	// Create a channel to collect errors
	errorChan := make(chan error, 6)

	// Fetch books
	wg.Add(1)
	go func() {
		defer wg.Done()
		books, err := GetTopBooksByUserID(db, userID)
		if err != nil {
			errorChan <- fmt.Errorf("books error: %v", err)
			return
		}
		mu.Lock()
		userData.Books = books
		mu.Unlock()
	}()

	// Fetch movies
	wg.Add(1)
	go func() {
		defer wg.Done()
		movies, err := getTopMoviesByUserID(db, userID)
		if err != nil {
			errorChan <- fmt.Errorf("movies error: %v", err)
			return
		}
		mu.Lock()
		userData.Movies = movies
		mu.Unlock()
	}()

	// Fetch pets
	wg.Add(1)
	go func() {
		defer wg.Done()
		pets, err := getTopPetByUserID(db, userID)
		if err != nil {
			errorChan <- fmt.Errorf("pets error: %v", err)
			return
		}
		mu.Lock()
		userData.Pets = pets
		mu.Unlock()
	}()

	// Fetch quotes
	wg.Add(1)
	go func() {
		defer wg.Done()
		quotes, err := getTopQuotesByUserID(db, userID)
		if err != nil {
			errorChan <- fmt.Errorf("quotes error: %v", err)
			return
		}
		mu.Lock()
		userData.Quotes = quotes
		mu.Unlock()
	}()

	// Fetch recipes
	wg.Add(1)
	go func() {
		defer wg.Done()
		recipes, err := getRecipeByUserID(db, userID)
		if err != nil {
			errorChan <- fmt.Errorf("recipes error: %v", err)
			return
		}
		mu.Lock()
		userData.Recipes = recipes
		mu.Unlock()
	}()

	// Fetch travel data
	wg.Add(1)
	go func() {
		defer wg.Done()
		travel, err := getTravelByUserID(db, userID)
		if err != nil {
			errorChan <- fmt.Errorf("travel error: %v", err)
			return
		}
		mu.Lock()
		userData.Travel = travel
		mu.Unlock()
	}()

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(errorChan)
	}()

	// Collect any errors
	for err := range errorChan {
		errors = append(errors, err)
	}

	// If there are critical errors, return them
	if len(errors) > 3 { // Allow some errors but not too many
		return nil, fmt.Errorf("too many data collection errors: %v", errors)
	}

	return userData, nil
}

// generatePersonalityAnalysis calls Gemini API to analyze personality
func generatePersonalityAnalysis(userData *UserData) (string, error) {
	// Get API key from environment
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY not found in environment variables")
	}

	// Create the prompt for personality analysis
	prompt := createPersonalityPrompt(userData)

	// Prepare Gemini API request
	geminiReq := GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: prompt},
				},
			},
		},
	}

	// Convert to JSON
	jsonData, err := json.Marshal(geminiReq)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Make API call to Gemini
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s", apiKey)
	
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Parse response
	var geminiResp GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content in API response")
	}

	return geminiResp.Candidates[0].Content.Parts[0].Text, nil
}

// createPersonalityPrompt creates a detailed prompt for personality analysis
func createPersonalityPrompt(userData *UserData) string {
	prompt := `As an expert personality analyst, please analyze the following user data and provide a comprehensive personality profile. Focus on personality traits, interests, values, and behavioral patterns.

User Data Analysis:

**Books:**
`
	for _, book := range userData.Books {
		if bookName, ok := book["book_name"].(string); ok {
			reason, _ := book["reason"].(string)
			prompt += fmt.Sprintf("- %s (Reason: %s)\n", bookName, reason)
		}
	}

	prompt += "\n**Movies/Shows:**\n"
	for _, movie := range userData.Movies {
		if title, ok := movie["title"].(string); ok {
			movieType, _ := movie["type"].(string)
			reason, _ := movie["reason"].(string)
			prompt += fmt.Sprintf("- %s (%s) - Reason: %s\n", title, movieType, reason)
		}
	}

	prompt += "\n**Pets:**\n"
	for _, pet := range userData.Pets {
		if name, ok := pet["name"].(string); ok {
			reason, _ := pet["reason"].(string)
			prompt += fmt.Sprintf("- %s (Reason: %s)\n", name, reason)
		}
	}

	prompt += "\n**Favorite Quotes:**\n"
	for _, quote := range userData.Quotes {
		if quoteText, ok := quote["quote"].(string); ok {
			prompt += fmt.Sprintf("- \"%s\"\n", quoteText)
		}
	}

	prompt += "\n**Recipes/Food:**\n"
	for _, recipe := range userData.Recipes {
		if name, ok := recipe["name"].(string); ok {
			reason, _ := recipe["reason"].(string)
			prompt += fmt.Sprintf("- %s (Reason: %s)\n", name, reason)
		}
	}

	prompt += "\n**Travel Destinations:**\n"
	for _, travel := range userData.Travel {
		if placeName, ok := travel["place_name"].(string); ok {
			reason, _ := travel["reason"].(string)
			prompt += fmt.Sprintf("- %s (Reason: %s)\n", placeName, reason)
		}
	}

	prompt += `

Please provide a detailed personality analysis covering:
1. **Core Personality Traits** - What kind of person they are
2. **Interests & Hobbies** - What they're passionate about
3. **Values & Beliefs** - What matters most to them
4. **Social Behavior** - How they interact with others
5. **Lifestyle Preferences** - How they like to live
6. **Growth Areas** - Potential areas for personal development
7. **Compatibility** - What kind of people/activities they'd connect with

Make the analysis personal, insightful, and constructive. Format it in a friendly, engaging way.`

	return prompt
}