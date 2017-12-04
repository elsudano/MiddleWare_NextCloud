package main

import (
	"log"
	"net/http"
    "time"
	"os"

	"github.com/elsudano/MiddleWare_NextCloud/webservice"
)

func main() {
	router := webservice.MyRouter()
	port := ":" + os.Getenv("PORT")
    server := &http.Server{
    	Addr:           port,
    	Handler:        router,
    	ReadTimeout:    10 * time.Second,
    	WriteTimeout:   10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(server.ListenAndServe())
}
