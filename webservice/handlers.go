package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/gorilla/mux"
)

// DEBUG es una constante que indica que se activa el modo depuración
// a lo largo de todo el codigo hay diferentes mensajes de pantalla
// para poder realizar un rustico depurado de la aplicación.
const DEBUG bool = true

// FStatus función que se encarga de responder a la disponibilidad que
// tiene nuestro webservice, siempre respondera de la misma manera,
// con un JSON con un campo status = OK.
func FStatus(w http.ResponseWriter, r *http.Request) {
	if DEBUG {
		fmt.Println("Acabamos de realizar la peticion del status")
	}
	status := Status{Status: "OK"}
	json.NewEncoder(w).Encode(status)
}

// FList funcion que se encarga de listar todos los eventos de NextCloud
func FList(w http.ResponseWriter, r *http.Request) {
	// Esta es la estructura de datos que vamos a enviar en formato
	// JSON al cliente, es una colección de objetos de JSON
	var responseJSON []EventJSON
	// Este es un evento de calandario en formato JSON
	var eventJSON EventJSON
	// con esto recorremos todo el XML
	// menos el primer Response que es donde estan los datos
	// del calendario al que accedemos
	for i, respTag := range xmlData.Responses {
		// en esta parte mostramos el resultado del JSON y formateamos la
		// salida por consola de un objeto tipo evento
		// evitamos el 0 por que son los datos del calendario completo
		if i != 0 && respTag.Propstat[0].Status == STATUSOK &&
		   strings.Contains(respTag.Propstat[0].Prop.Content_Type, "component=vevent") {
			eventJSON = EventJSON {
				// Con esta linea eliminamos todo lo que sobra para quedarnos
				// solo con el ID del evento
				Id: strings.TrimSuffix(strings.TrimPrefix(respTag.Href, URL_BASE+"Nextcloud-"), ".ics"),
				// Lo que hacemos con replace es quitar las comillas que molestan
				// en el campo ETAG que vienen por defecto en el XML
				Etag:     strings.Replace(respTag.Propstat[0].Prop.Etag, "\"", "", -1),
				Modified: respTag.Propstat[0].Prop.Modified,
			}
			// añadimos el evento a la respuesta que vamos a dar en JSON
			responseJSON = append(responseJSON, eventJSON)
		}
	}
	// Con esto rellenamos el JSON que vamos a devolver al cliente
	if err := json.NewEncoder(w).Encode(responseJSON); err != nil {
		if DEBUG {
			fmt.Println("Este error es del JSON")
		}
		panic(err)
	}
	if DEBUG {
		printXML(xmlData)
	}
}

// se llama cuando se accede a "/show/{id}"
func FShow(w http.ResponseWriter, r *http.Request) {
	// recogemos las variables de la URI
	vars := mux.Vars(r)
	Id := vars["id"]
	// Este es un evento de calandario en formato JSON
	var eventJSON icsJSON

	if DEBUG {
		fmt.Println("Estas son las variables de entrada:", vars)
	}
	// con esto recorremos todo el XML
	for i , respTag := range xmlData.Responses {
		if i != 0 && strings.TrimSuffix(strings.TrimPrefix(respTag.Href, URL_BASE+"Nextcloud-"), ".ics") == Id {
			stringICS := readICS(eventJSON, URL_COMPLETE_PATH + "Nextcloud-" + Id + ".ics")
			if DEBUG {
				fmt.Println("Este es el contenido del fichero ICS\n", stringICS)
			}
		}
	}
	// Con esto rellenamos el JSON que vamos a devolver al cliente
	if err := json.NewEncoder(w).Encode(eventJSON); err != nil {
		if DEBUG {
			fmt.Println("Este error es del JSON")
		}
		panic(err)
	}
}

// se llama cuando se accede a "/new"
func FNew(w http.ResponseWriter, r *http.Request) {

}

// se llama cuando se accede a "/update/{id}"
func FUpdate(w http.ResponseWriter, r *http.Request) {

}

// se llama cuando se accede a "/delete/{id}"
func FDelete(w http.ResponseWriter, r *http.Request) {

}

// se llama cuando se accede a "/refresh"
// Se usa para refrescar la base de datos temporal
func FRefresh(w http.ResponseWriter, r *http.Request) {
	xmlData = readXML()
}
// se llama cuando se accede a "/exit"
func FExit(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
