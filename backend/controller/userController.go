package controller

import (
	"github.com/ekharisma/poltekkes-webservice/entity"
	"github.com/ekharisma/poltekkes-webservice/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUserController interface {
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
	GetUsers(c *gin.Context)
}

type UserController struct {
	UserModel model.IUserModel
}

func CreateUserController(model model.IUserModel) IUserController {
	return &UserController{UserModel: model}
}

func (u UserController) CreateUser(c *gin.Context) {
	var payload entity.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := u.UserModel.Create(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}

func (u UserController) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := u.UserModel.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u UserController) GetUsers(c *gin.Context) {
	users, err := u.UserModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if users == nil {
		users = make([]*entity.User, 0)
	}
	c.JSON(http.StatusOK, users)
}
