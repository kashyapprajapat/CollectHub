package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name        string             `json:"name" bson:"name"`
    Ingredients string             `json:"ingredients" bson:"ingredients"`
    Reason      string             `json:"reason" bson:"reason"`
    UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
}
