package models

import (
	"fmt"
)

// Usuario con todos sus datos
type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
	Img      []byte `json:"img"`
}

func (user User) String() string {
	return fmt.Sprintf("Id: %s, Username: %s, Email: %s, Pass: %s, Img is ommitted", user.Id, user.Username, user.Email, user.Pass)
}
