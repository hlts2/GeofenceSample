package main

import (
	"log"

	"github.com/hlts2/GeofenceSample/GeofenceManageServer/server"
)

func main() {
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}
