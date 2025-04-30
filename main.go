package main

import (
	"database/sql"
	"log"

	"dotted_backend/controller"
	"dotted_backend/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func init() {
	db = controller.NewConnection()
}

func main() {
	router := gin.Default()

	router.POST("/api/users", func(c *gin.Context) {
		var newUser models.User

		if err := c.BindJSON(&newUser); err != nil {
			log.Println("Error while binding JSON: ", err.Error())
		}

		log.Println(newUser)

		res, _ := controller.NewUser(db, newUser)

		if res == -1 {
			c.JSON(409, gin.H{"msg": "Username or email taken!."})
		} else if res == 0 {
			c.JSON(500, gin.H{"msg": "Error, the user couldn't be added."})
		} else {
			c.JSON(201, gin.H{"msg": "User added!"})
		}
	})

	router.GET("/api/store-user-data", controller.CORSMiddleware())

	router.Run(":8000")
}
