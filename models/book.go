package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    BookName    string             `bson:"book_name" json:"book_name"`
    Author      string             `bson:"author" json:"author"`
    Reason      string             `bson:"reason" json:"reason"`
    UserID      primitive.ObjectID `bson:"user_id" json:"user_id"` // Reference to User
}
