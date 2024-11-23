package models

import "time"

type Message struct {
	OriginalMessage  string    `bson:"original_message"`
	ProcessedMessage string    `bson:"processed_message"`
	Timestamp        time.Time `bson:"timestamp"`
}
