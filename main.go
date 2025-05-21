package main

import (
	"database/sql"
	"dotted_backend/controller"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func init() {
	db = controller.NewConnection()
}

func main() {
	router := gin.Default()

	router.POST("/api/users", controller.AddNewUserGinHandler(db))
	router.GET("/api/users/:userId", controller.GetUserGinHandler(db))
	// maybe later router.GET("/api/user-posts/:userId", controller.GetUserPostsGinHandler(db))

	router.POST("/api/posts", controller.AddUserPostGinHandler(db))

	router.GET("/api/store-user-data", controller.CORSMiddleware())

	router.Run(":8000")
}
