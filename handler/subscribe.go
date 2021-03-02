package handler

import (
	"lyra/subscribe"

	"github.com/gin-gonic/gin"
)

type subscribeHandler struct {
	subscribeService subscribe.Service
}

func SubscribeHandler(subscribe subscribe.Service) *subscribeHandler {
	return &subscribeHandler{subscribe}
}

func (h *subscribeHandler) Create(c *gin.Context) {
	var input subscribe.SubscribeInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}

	newSubscribe, err := h.subscribeService.Create(input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	c.JSON(400, gin.H{
		"status": "success",
		"data":   subscribe.SubscribeFormat(newSubscribe),
	})
	return
}

func (h *subscribeHandler) FindAll(c *gin.Context) {
	subscribe, _ := h.subscribeService.Find()
	c.JSON(400, gin.H{
		"status": "success",
		"data":   subscribe,
	})
	return

}

func (h *subscribeHandler) Delete(c *gin.Context) {
	var input subscribe.SubscribeDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	_, err = h.subscribeService.DeleteSubscribe(input)
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "data deleted",
	})
	return
}
