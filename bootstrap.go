package main

import (
	"log"

	"github.com/christian-gama/uow/sql"
	"github.com/christian-gama/uow/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DB is a global variable holding the database connection instance.
var DB *gorm.DB

func init() {
	DB = sql.Conn()

	user.Migrate(DB)
}

// Bootstrap is a function that sets up the application by creating the user repository, user service, and user controller.
// It then registers the necessary routes to the provided server instance, runs the server, and listens for incoming requests.
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
	server.POST("/users/transfer", userController.TransferMoney)
	server.GET("/users/:id", userController.FindOne)

	log.Println("Application running on port 8080")
	server.Run(":8080")
}
