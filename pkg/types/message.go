package types

import "time"

// Message defines the structure for our JSON messages.
type Message struct {
	Sender    string    `json:"sender"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

