package models

type UserPosts struct {
	User  User   `json:"user"`
	Posts []Post `json:"posts"`
}
