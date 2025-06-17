package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quote struct {
    ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Quote  string             `json:"quote" bson:"quote"`
    Author string             `json:"author" bson:"author"`
    UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
}
