package webservice

import (
	//"fmt"
)

// Estas son las constantes que usaremos en el programa para comprobar
// los diferentes estados en los que se encuentra el flujo de la
// aplicación
const (
    // cuando el estado de la respuesta http es correcto
    STATUSOK = "HTTP/1.1 200 OK"
    // cuando el estado de la respuesta http es no encontrado
    STATUSNOTFOUND = "HTTP/1.1 404 Not Found"
)
