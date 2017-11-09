package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

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
