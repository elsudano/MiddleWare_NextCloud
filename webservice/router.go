package webservice

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MyRouter es la función principal del programa, se utiliza para realizar
// una abstracción de todo lo que conlleva una comunicación http(s).
//
// Con esta función lo que hacemos es recorrer la lista de posibilidades
// que tenemos en nuestro webservice y preparar nuestra aplicación
// para contestar a cada una de ellas de una manera muy concreta.
//
// Esa manera de contestar se define en las diferentes funciones que hacemos
// implementado para cada una de las URLs a las que queremos dar respuesta en
// nuestro webservice.
func MyRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
