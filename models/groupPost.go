package models

import "time"

type GroupPost struct {
	PostId     string `json:"postId"`
	GroupId    string `json:"groupId"`
	UploaderId string `json:"uploaderId"`
	// must be video, audio, text or image
	Type    string    `json:"type"`
	Value   []byte    `json:"value"`
	PubTime time.Time `json:"pubTime"`
}
