package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Health(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
	})
}

func (h *handler) GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
