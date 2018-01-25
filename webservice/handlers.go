package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
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

// FShow Función que se encarga de mostrar el evento que se elecciona
// por el ID del mismo.
//
// Forma de uso:
// URL/show/{ID} donde el parametro es uno de entrada y tiene que
// coresponder con uno de los que se muestran en la función listar.
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
			// guardamos el evento en su structura que devolvemos
			eventJSON = readICS(URL_COMPLETE_PATH + "Nextcloud-" + Id + ".ics")
			// esto lo hacemos para mantener el mismo Id que sale en la lista
			// en NextCloud se utilizan diferentes IDs para gestionar mas cosas
			eventJSON.Id = Id
			if DEBUG {
				fmt.Println("Este es el contenido del fichero ICS\n", eventJSON)
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
	// Este es un evento de calandario en formato JSON
	var eventJSON icsJSON
	// Con esto leemos los datos recogidos desde el cliente
	if err := json.NewDecoder(r.Body).Decode(&eventJSON); err != nil {
		if DEBUG {
			fmt.Println("Este error es del JSON")
		}
		panic(err)
	}
	eventJSON.Id = "KJ43MS30MEJQIR6ULBI0U9"
	fmt.Println(eventJSON)
	// BEGIN:VCALENDAR
	// PRODID:-//Nextcloud calendar v1.5.6
	// VERSION:2.0
	// CALSCALE:GREGORIAN
	// BEGIN:VEVENT
	// CREATED:20171112T165623
	// DTSTAMP:20171112T165623
	// LAST-MODIFIED:20171206T172154
	// UID:JZ100KIU23DVQVZIVMWG5
	// SUMMARY:Esto es una prueba de concepto
	// CLASS:PUBLIC
	// STATUS:CONFIRMED
	// DTSTART;TZID=Africa/Ceuta:20171207T170000
	// DTEND;TZID=Africa/Ceuta:20171207T180000
	// END:VEVENT
	// BEGIN:VTIMEZONE
	// TZID:Africa/Ceuta
	// BEGIN:DAYLIGHT
	// TZOFFSETFROM:+0100
	// TZOFFSETTO:+0200
	// TZNAME:CEST
	// DTSTART:19700329T020000
	// RRULE:FREQ=YEARLY;BYMONTH=3;BYDAY=-1SU
	// END:DAYLIGHT
	// BEGIN:STANDARD
	// TZOFFSETFROM:+0200
	// TZOFFSETTO:+0100
	// TZNAME:CET
	// DTSTART:19701025T030000
	// RRULE:FREQ=YEARLY;BYMONTH=10;BYDAY=-1SU
	// END:STANDARD
	// END:VTIMEZONE
	// END:VCALENDAR
	// // datos en bruto del fichero subido
	// bodyRaw, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// if DEBUG {
	// 	fmt.Println("Los datos de entrada son:")
	// 	fmt.Printf("\tBody: %s\n", bodyRaw)
	// }
	defer r.Body.Close()
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

// se llama cuando se accede a "/help"
// Muestra la información relativa al webservice
// Toda la funcionalidad que tiene y que metodos se pueden usar
func FHelp(w http.ResponseWriter, r *http.Request) {
	var data []byte
	var err error
	if os.Getenv("HOME") == "/app" {
		data, err = ioutil.ReadFile("/app/webservice/wshelp.html")
	} else {
		data, err = ioutil.ReadFile(GOPATH + "/src/github.com/elsudano/MiddleWare_NextCloud/webservice/wshelp.html")
	}

    if err == nil {
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Charset", "utf-8")
		w.Write(data)
	} else {
		if DEBUG {
			fmt.Println("Error en la lectura:", err)
		}
	    w.WriteHeader(404)
	    w.Write([]byte(http.StatusText(404)))
	}
}

// se llama cuando se accede a "/exit"
func FExit(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
