package handler

import (
	"lyra/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func UserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var input user.Register

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}

	userNew, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"data":   user.UserFormat(userNew),
	})
	return
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.Login
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(422, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	logginUser, err := h.userService.Login(input)
	if err != nil {
		c.JSON(422, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"data":   user.UserFormat(logginUser),
	})
	return
}

func (h *userHandler) TokenUser(c *gin.Context) {
	var input user.TokenInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	newToken, err := h.userService.Token(input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"data":   user.TokenFormate(newToken),
	})
	return
}
