package models

import "time"

type URL struct {
	ID string `bson:"_id, omitempty"`
	LongURL string `bson:"long_url"`
	CreateAt time.Time `bson:"created_at"`
}