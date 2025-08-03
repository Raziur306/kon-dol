package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Incident struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
	Title     string             `bson:"title" json:"title"`
	Location  string             `bson:"location" json:"location"`
	Party     Party              `bson:"party" json:"party"`
	Date      string             `bson:"date" json:"date"`
	Source    string             `bson:"source" json:"source"`
	ShortDesc string             `bson:"short_desc" json:"short_desc"`
	Link      string             `bson:"link" json:"link"`
	TrackId   string             `bson:"track_id" json:"track_id"`
}
