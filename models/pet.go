package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pet struct {
    ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name   string             `json:"name" bson:"name"`
    Reason string             `json:"reason" bson:"reason"`
    UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
}
