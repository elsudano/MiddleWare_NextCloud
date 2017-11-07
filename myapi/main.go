package main

import (
	"log"
	"net/http"
    "time"
)

func main() {

	router := MyRouter()

    server := &http.Server{
    	Addr:           ":5000",
    	Handler:        router,
    	ReadTimeout:    10 * time.Second,
    	WriteTimeout:   10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(server.ListenAndServe())
}
