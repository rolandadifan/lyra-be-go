package handler

import (
	"lyra/request"

	"github.com/gin-gonic/gin"
)

type requestHandler struct {
	requestService request.Service
}

func RequestHandler(requestService request.Service) *requestHandler {
	return &requestHandler{requestService}
}

func (h *requestHandler) Create(c *gin.Context) {
	var input request.RequestInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}

	newRequest, err := h.requestService.CreateData(input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	c.JSON(400, gin.H{
		"status": "success",
		"data":   newRequest,
	})
	return
}

func (h *requestHandler) Find(c *gin.Context) {
	request, err := h.requestService.FindData()
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	c.JSON(400, gin.H{
		"status": "success",
		"data":   request,
	})
	return
}

func (h *requestHandler) Delete(c *gin.Context) {
	var input request.RequestDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	_, err = h.requestService.DestroyData(input)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "data deleted",
	})
	return

}
