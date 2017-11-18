package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index Add Index Handler
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
