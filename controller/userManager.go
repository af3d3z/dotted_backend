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

	stmtInsert, err := db.Prepare("INSERT INTO users (id, email, username, img) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Println("Couldn't prepare the user for insertion:", err)
		return 0, err
	}
	defer stmtInsert.Close()

	res, err := stmtInsert.Exec(user.Id, user.Email, user.Username, user.Img)
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

func GetUser(db *sql.DB, userId string) (models.User, error) {
	var user models.User

	var img []byte
	var description sql.NullString

	err := db.QueryRow("SELECT id, username, email, img, description FROM users where id = ?", userId).Scan(&user.Id, &user.Username, &user.Email, &img, &description)
	if err != nil {
		log.Println("Couldn't access user ", userId, user.Email, err.Error())
	}

	if img != nil {
		user.Img = make([]byte, len(img))
		copy(user.Img, img)
	} else {
		user.Img = nil
	}

	if description.Valid {
		user.Description = description.String
	} else {
		user.Description = ""
	}

	return user, err
}

func GetPosts(db *sql.DB, userId string) ([]models.Post, error) {
	var posts []models.Post

	rows, err := db.Query("SELECT * FROM posts WHERE userId = ?", userId)
	if err != nil {
		log.Println(err.Error())
	} else {
		for rows.Next() {
			var post models.Post

			err := rows.Scan(&post.PostId, &post.UserId, &post.PubTime, &post.Type, &post.Value)
			if err != nil {
				log.Println(err.Error())
			}

			posts = append(posts, post)
		}
	}

	return posts, err
}

func AddPost(db *sql.DB, post models.Post) bool {
	inserted := true
	query := "INSERT INTO posts (userId, pubTime, type, value) VALUES (?, ?, ?, ?)"
	insert, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		inserted = false
	} else {
		res, err := insert.Exec(post.UserId, post.PubTime, post.Type, post.Value)
		if err != nil {
			log.Println(err.Error())
			log.Println(post.UserId, post.PubTime, post.Type)
			inserted = false
		} else {
			insert.Close()
			rows, err := res.RowsAffected()
			if err != nil {
				log.Println(err.Error())
				inserted = false
			} else {
				if rows != 1 {
					log.Println("Could not insert")
					inserted = false
				}
			}
		}
	}

	return inserted

}
