package controller

import (
	"database/sql"
	"dotted_backend/models"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// contains the handler that creates a new user
func AddNewUserGinHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User

		if err := c.BindJSON(&newUser); err != nil {
			log.Println("Error while binding JSON: ", err.Error())
		}

		log.Println(newUser)

		res, _ := NewUser(db, newUser)

		if res == -1 {
			c.JSON(409, gin.H{"msg": "Username or email taken!."})
		} else if res == 0 {
			c.JSON(500, gin.H{"msg": "Error, the user couldn't be added."})
		} else {
			c.JSON(201, gin.H{"msg": "User added!"})
		}
	}
}

// returns the user info stored in the db
func GetUserGinHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		userId := c.Param("userId")

		user, err := GetUser(db, userId)
		if err != nil {
			c.JSON(500, gin.H{"msg": "Error: " + err.Error()})
		}

		c.JSON(200, user)
	}
}

// returns a user with all their posts
func GetUserPostsGinHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userPosts models.UserPosts

		userId := c.Param("userId")

		user, err := GetUser(db, userId)
		if err != nil {
			c.JSON(500, gin.H{"msg": "Error: " + err.Error()})
		} else {
			posts, err := GetPosts(db, userId)
			if err != nil {
				c.JSON(500, gin.H{"msg": "Error: " + err.Error()})
			}

			userPosts.User = user
			userPosts.Posts = posts

			c.JSON(200, gin.H{"userPosts": userPosts})
		}
	}
}

func AddUserPostGinHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileValue, err := c.FormFile("value")
		if err != nil {
			c.JSON(400, gin.H{"msg": "Couldn't upload the file."})
		}

		file, err := fileValue.Open()
		if err != nil {
			c.JSON(500, gin.H{"msg": "File unreadable."})
		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(500, gin.H{"msg": "Error while reading file"})
		}

		post := models.Post{
			UserId:  c.PostForm("userId"),
			Type:    c.PostForm("type"),
			PubTime: time.Now(),
			Value:   fileBytes,
		}

		if AddPost(db, post) {
			c.JSON(201, gin.H{"msg": "Uploaded successfully"})
		} else {
			c.JSON(415, gin.H{"msg": "Could not upload."})
		}

	}
}
