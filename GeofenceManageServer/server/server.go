package server

import (
	"github.com/gin-gonic/gin"
)

// Init -- Server Initialization
func Init() error {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

	}
	return nil
}
