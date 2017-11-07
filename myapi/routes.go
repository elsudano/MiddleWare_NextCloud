package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func MyRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{
        "Exit",
        "GET",
        "/exit",
        Exit,
    },
    Route{
        "Root",
        "GET",
        "/",
        Root,
    },
    Route{
        "Index",
        "GET",
        "/index",
        Index,
    },
    Route{
        "Show",
        "GET",
        "/{todoId}",
        Show,
    },
}
