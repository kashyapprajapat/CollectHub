package controllers

import (
    "context"
    "github.com/kashyapprajapat/collecthub_api/models"
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitUserController(db *mongo.Database) {
    userCollection = db.Collection("users")
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPasswordHash compares password with hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CreateUser(c *fiber.Ctx) error {

    var user models.User
    if err := c.BodyParser(&user); err != nil {
        log.Printf("Body parsing error: %v", err)
        return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    
    // Validate required fields
    if user.Name == "" || user.Email == "" || user.Password == "" {
        return c.Status(400).JSON(fiber.Map{"error": "name, email, and password are required"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Check if user with email already exists
    var existingUser models.User
    err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
    if err == nil {
        return c.Status(409).JSON(fiber.Map{"error": "user with this email already exists"})
    }

    // Hash the password
    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to hash password"})
    }
    user.Password = hashedPassword

    // Generate new ObjectID for the user
    user.ID = primitive.NewObjectID()

    res, err := userCollection.InsertOne(ctx, user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to insert user"})
    }

    // Return user without password
    response := fiber.Map{
        "id":      res.InsertedID,
        "name":    user.Name,
        "email":   user.Email,
        "message": "user created successfully",
    }

    return c.Status(201).JSON(response)
}

func GetUsers(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := userCollection.Find(ctx, bson.M{})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to fetch users"})
    }
    defer cursor.Close(ctx)
    
    var users []models.User
    if err = cursor.All(ctx, &users); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to decode users"})
    }

    // Remove passwords from response (extra safety)
    var safeUsers []fiber.Map
    for _, user := range users {
        safeUsers = append(safeUsers, fiber.Map{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
        })
    }

    return c.JSON(safeUsers)
}

func LoginUser(c *fiber.Ctx) error {
    var loginReq models.LoginRequest
    if err := c.BodyParser(&loginReq); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    // Validate required fields
    if loginReq.Email == "" || loginReq.Password == "" {
        return c.Status(400).JSON(fiber.Map{"error": "email and password are required"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Find user by email
    var user models.User
    err := userCollection.FindOne(ctx, bson.M{"email": loginReq.Email}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(401).JSON(fiber.Map{"error": "invalid email or password"})
        }
        return c.Status(500).JSON(fiber.Map{"error": "failed to find user"})
    }

    // Check password
    if !CheckPasswordHash(loginReq.Password, user.Password) {
        return c.Status(401).JSON(fiber.Map{"error": "invalid email or password"})
    }

    // Login successful - return user data without password
    response := models.LoginResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
        // Token: "your-jwt-token-here", // Add JWT token if needed
    }

    return c.JSON(fiber.Map{
        "message": "login successful",
        "user":    response,
    })
}