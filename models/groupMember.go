package models

import "time"

type GroupMember struct {
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
	// must be admin or member
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joinedAt"`
}
