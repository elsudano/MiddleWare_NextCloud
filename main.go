package main

import (
	"log"
	"net/http"
    "time"
	"os"
	"fmt"

	"github.com/elsudano/MiddleWare_NextCloud/webservice"
)

const DEBUG bool = true

func main() {
	fmt.Println(DEBUG)
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
