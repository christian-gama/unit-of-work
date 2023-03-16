package main

import (
	"log"

	"github.com/christian-gama/uow/sql"
	"github.com/christian-gama/uow/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = sql.Conn()

	user.Migrate(DB)
}

func Bootstrap(server *gin.Engine) {
	log.Println("Bootstrapping the application")

	log.Println("Creating the user repository")
	userRepo := user.NewRepository(sql.NewRepository())

	log.Println("Creating the user service")
	userService := user.NewService(userRepo, sql.NewUnitOfWork(DB))

	log.Println("Creating the user controller")
	userController := user.NewController(userService)

	server.GET("/users", userController.FindAll)
	server.POST("/users", userController.Save)

	log.Println("Application running on port 8080")
	server.Run(":8080")
}
