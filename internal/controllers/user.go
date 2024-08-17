package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/me2seeks/cola/internal/services"
	"github.com/me2seeks/cola/internal/types"
)

var userService *services.UserService

type UserController struct{}

func NewUserController() *UserController {
	userService = services.NewUserService()
	return &UserController{}
}

func (u *UserController) Login(c *gin.Context) {
	var req *types.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := userService.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{})
}

func (u *UserController) Register(c *gin.Context) {
	var req *types.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := userService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{})
}

func (u *UserController) GetUser(c *gin.Context) {
	user, err := userService.GetUser(c.GetString("ID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var req *types.UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userService.UpdateUser(c.GetString("ID"), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	err := userService.DeleteUser(c.GetString("ID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (u *UserController) ListUser(c *gin.Context) {
	users, err := userService.ListUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (u *UserController) ListUserByPage(c *gin.Context) {
	var req *types.ListUserByPageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, total, err := userService.ListUserByPageWithTotal(req.PageNum, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users, "total": total})
}
