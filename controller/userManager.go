package controller

import (
	"database/sql"
	"log"

	"dotted_backend/models"
)

func NewUser(db *sql.DB, user models.User) (int64, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? OR username = ?", user.Email, user.Username).Scan(&count)
	if err != nil {
		log.Println("Error checking for existing user:", err)
		return 0, err
	}
	if count > 0 {
		log.Println("User with this email or username already exists.")
		return -1, nil
	}

	stmtInsert, err := db.Prepare("INSERT INTO users (id, email, username, img) VALUES(uuid(), ?, ?, ?)")
	if err != nil {
		log.Println("Couldn't prepare the user for insertion:", err)
		return 0, err
	}
	defer stmtInsert.Close()

	res, err := stmtInsert.Exec(user.Email, user.Username, user.Img)
	if err != nil {
		log.Println("Couldn't insert the user:", err)
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("Error when accessing affected rows:", err)
		return 0, err
	}

	return rows, nil
}
