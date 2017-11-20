package top

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TopHandler Add Top Handler
func TopHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
