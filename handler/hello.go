package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (handler *HelloHandler) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
