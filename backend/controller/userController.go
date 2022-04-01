package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ekharisma/poltekkes-webservice/entity"
	"github.com/ekharisma/poltekkes-webservice/model"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
	GetUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
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
	fmt.Println(len(users))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(users) == 0 {
		fmt.Println("Hasil query 0")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User Table is empty",
		})
		return
	}
	if users == nil {
		users = make([]*entity.User, 0)
	}
	c.JSON(http.StatusOK, users)
}

func (u UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	var payload entity.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Wrong payload format",
		})
		return
	}
	if err := u.UserModel.Update(id, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Status(http.StatusAccepted)
}

func (u UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Bad Request",
		})
		return
	}
	if err := u.UserModel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Status(http.StatusAccepted)
}
