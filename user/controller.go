package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller is an interface that defines the functions necessary for a user controller.
type Controller interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	TransferMoney(ctx *gin.Context)
	FindOne(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
	users, err := c.service.FindAll(ctx.Request.Context(), &FindAllParams{})
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
	var dto *SaveDto
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saveParams := &SaveParams{
		&User{
			ID:    dto.ID,
			Name:  dto.Name,
			Money: dto.Money,
			Age:   dto.Age,
		},
	}
	err = c.service.Save(ctx.Request.Context(), saveParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, saveParams)
}

// TransferMoney is a function that transfers money from one user to another.
// It uses the user service instance to call the appropriate function and transfers money from one user to another.
// It then returns a status OK to the client.
func (c *controllerImpl) TransferMoney(ctx *gin.Context) {
	var dto *TransferMoneyDto
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transferMoneyParams := &TransferMoneyParams{
		From:   dto.From,
		To:     dto.To,
		Amount: dto.Amount,
	}
	err = c.service.TransferMoney(ctx.Request.Context(), transferMoneyParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "money transferred successfully"})
}

// FindOne is a function that returns a single user from the database.
// It uses the user service instance to call the appropriate function and returns the result.
func (c *controllerImpl) FindOne(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findOneParams := &FindOneParams{
		ID: uint(id),
	}
	user, err := c.service.FindOne(ctx.Request.Context(), findOneParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Delete is a function that deletes a user from the database.
// It uses the user service instance to call the appropriate function and deletes the user from the database.
// It then returns a status OK to the client.
func (c *controllerImpl) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteParams := &DeleteParams{
		ID: uint(id),
	}
	err = c.service.Delete(ctx.Request.Context(), deleteParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
