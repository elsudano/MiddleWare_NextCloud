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
        FExit,
    },
    Route{
        "Status",
        "GET",
        "/",
        FStatus,
    },
    Route{
        "Status",
        "GET",
        "/status",
        FStatus,
    },
    Route{
        "List",
        "GET",
        "/list",
        FList,
    },
    Route{
        "Show",
        "GET",
        "/show/{id}",
        FShow,
    },
    Route{
        "New",
        "POST",
        "/new",
        FNew,
    },
    Route{
        "Update",
        "POST",
        "/update/{id}",
        FUpdate,
    },
    Route{
        "Delete",
        "GET",
        "/delete/{id}",
        FDelete,
    },
}
