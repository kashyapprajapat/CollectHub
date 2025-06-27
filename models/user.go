package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name     string             `bson:"name" json:"name"`
    Email    string             `bson:"email" json:"email"`
    Password string             `bson:"password" json:"password,omitempty"` // Allow parsing but can be omitted in responses
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    ID    primitive.ObjectID `json:"id"`
    Name  string             `json:"name"`
    Email string             `json:"email"`
    Token string             `json:"token,omitempty"` // Optional: for JWT token
}