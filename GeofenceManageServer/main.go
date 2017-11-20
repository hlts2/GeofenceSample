package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hlts2/GeofenceSample/GeofenceManageServer/application/handler"
	"github.com/hlts2/GeofenceSample/GeofenceManageServer/application/handler/location"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("ui/template/*.html")
	r.GET("/", handler.RootHandler)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/locations", location.GetLocationsHandler)
		v1.GET("/locations/:id", location.GetLocationHandler)
		v1.POST("/locations", location.PostLocationHandler)
	}

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
