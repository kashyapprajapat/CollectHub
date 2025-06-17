package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type TravelBuddy struct {
    ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    PlaceName    string             `json:"place_name" bson:"place_name"`
    DateVisited  time.Time          `json:"date_visited" bson:"date_visited"`
    Reason       string             `json:"reason" bson:"reason"`
    UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
}
