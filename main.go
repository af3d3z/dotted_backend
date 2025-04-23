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

	router.POST("/users", func(c *gin.Context) {
		var newUser models.User

		if err := c.BindJSON(&newUser); err != nil {
			log.Println("Error while binding JSON: ", err.Error())
		}

		res := controller.NewUser(db, newUser)

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Println("Error while getting the rows result: ", err.Error())
		}

		if rowsAffected != 1 {
			c.JSON(500, "Error, the user couldn't be added.")
		} else {
			c.JSON(200, "User added!")
		}

	})

	router.Run(":8000")
}
