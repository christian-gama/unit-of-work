package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type controllerImpl struct {
	service Service
}

func (c *controllerImpl) FindAll(ctx *gin.Context) {

	users, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

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

func NewController(service Service) Controller {
	return &controllerImpl{service}
}
