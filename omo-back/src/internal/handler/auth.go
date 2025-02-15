package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"omo-back/src/internal/service/impl"
)

type AuthHandler struct {
	userService *impl.UserServiceImpl
}

func NewAuthHandler(userService *impl.UserServiceImpl) *AuthHandler {
	return &AuthHandler{userService: userService}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthHandler) Register(c *gin.Context) {
	request := RegisterRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := a.userService.Register(c, request.Username, request.Email, request.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":  "user created",
		"username": request.Username,
		"email":    request.Email,
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}
