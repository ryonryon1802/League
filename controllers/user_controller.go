package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/models/implements"
	"log"
	"net/http"
)

var (
	UserController IUserController
)

func init() {
	UserController = NewUserController(implements.UserImplements)
}

type userController struct {
	implements.IUserImplements
}

type IUserController interface {
	IndexUserController(c *gin.Context)
	IndexOneUserController(c *gin.Context)
	CreateUserController(c *gin.Context)
	UpdateUserController(c *gin.Context)
	DestroyUserController(c *gin.Context)
}

func NewUserController(impl implements.IUserImplements) IUserController {
	return &userController{impl}
}

func (ctrl *userController) IndexUserController(c *gin.Context) {
	users, err := ctrl.IUserImplements.SelectAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *userController) IndexOneUserController(c *gin.Context) {
	users, err := ctrl.IUserImplements.SelectSingleUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *userController) CreateUserController(c *gin.Context) {
	err := ctrl.IUserImplements.CreateUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl *userController) UpdateUserController(c *gin.Context) {
	err := ctrl.IUserImplements.UpdateUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl *userController) DestroyUserController(c *gin.Context) {
	err := ctrl.IUserImplements.DeleteUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}