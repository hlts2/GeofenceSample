package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootHandler Add Root Handler
func RootHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
