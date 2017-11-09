package main

import (
	"encoding/json"
	"fmt"
    "os"
	"net/http"

	"github.com/gorilla/mux"
)

func Root(w http.ResponseWriter, r *http.Request) {
    status := Status{Status: "OK"}
    json.NewEncoder(w).Encode(status)
}

func Index(w http.ResponseWriter, r *http.Request) {
    objects := Objects{
        Object{Name: "Write presentation"},
        Object{Name: "Host meetup"},
    }
    if err := json.NewEncoder(w).Encode(objects); err != nil {
        panic(err)
    }
}

func Show(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Exit(w http.ResponseWriter, r *http.Request) {
    os.Exit(0)
}
