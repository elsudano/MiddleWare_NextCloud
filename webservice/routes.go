package webservice

import (
    "net/http"
)

// Route estructura que se encarga de asociar los diferentes datos de una petición.
//
// El nombre es el valor que le podemos dar a nuestra petición.
//
// El metodo es de que manera se realiza la llamada a la URL.
//
// El patrón es la dirección URL a la que se contesta.
//
// El manejador es la función a la que hace referencia esa URL.
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes es un conjunto de Route el cual almacena las diferentes
// opciones que tenemos.
type Routes []Route

// routes definición de todas las opciones que dispone el webservice.
//
// Exit se utiliza para pruebas, cierra el webservice.
//
// Status en ambos casos devuelve un JSON con el campo status = ok
// sirve para poder saber si el webservice se encuentra operativo.
//
// List muestra una lista de los diferentes items que a los que
// podemos acceder a través de la API.
//
// Show realiza una solicitud de uno de los items, con el id
// especificado.
//
// New genera un item nuevo y lo añade a NextCloud.
//
// Update actualiza un item que se especifica con el id.
//
// Delete borra un item que se especifica con el id.
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
