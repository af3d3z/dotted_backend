package models

import (
	"time"
)

type Post struct {
	PostId  string    `json:"postId"`
	UserId  string    `json:"userId"`
	PubTime time.Time `json:"pubTime"`
	Type    string    `json:"type"`
	Value   []byte    `json:"value"`
}
