package controller

import (
	"database/sql"
	"log"

	"dotted_backend/models"
)

func NewUser(db *sql.DB, user models.User) sql.Result {
	stmtInsert, err := db.Prepare("INSERT INTO users VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Couldn't prepare the user for insertion: ", err.Error())
	}
	defer stmtInsert.Close()

	res, err := stmtInsert.Exec(user.Id, user.Username, user.Email, user.Pass, user.Img)
	if err != nil {
		log.Println("Couldn't insert the user:", err.Error())
	}

	return res
}
