package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
    ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title  string             `json:"title" bson:"title"`
    Type   string             `json:"type" bson:"type"` // e.g., "movie", "series"
    Reason string             `json:"reason" bson:"reason"`
    UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
}
