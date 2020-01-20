package main

import (
	"log"
	"translate/router"
	"translate/server"
)


func main() {

	mux := router.SetupRouter()
	srv := server.New(mux, ":8080")
	log.Fatalf("server failed to start: %v", srv.ListenAndServe())
}
