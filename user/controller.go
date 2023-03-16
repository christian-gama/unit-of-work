package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is an interface that defines the functions necessary for a user controller.
type Controller interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type controllerImpl struct {
	service Service
}

// NewController is a function that creates a new user controller instance using the provided user service.
func NewController(service Service) Controller {
	return &controllerImpl{service}
}

// FindAll is a function that returns all the users in the database.
// It uses the user service instance to call the appropriate function and returns the result.
func (c *controllerImpl) FindAll(ctx *gin.Context) {
	users, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// Save is a function that creates a new user in the database.
// It uses the user service instance to call the appropriate function and saves the user to the database.
// It then returns the created user and a status OK to the client.
func (c *controllerImpl) Save(ctx *gin.Context) {
	var user User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.service.Save(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
