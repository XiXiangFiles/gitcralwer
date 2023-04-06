package model

import "time"

type Commitment struct {
	Author    string    `json:"author"`
	Comment   string    `json:"comment"`
	Timestamp time.Time `json:"timestamp"`
}
