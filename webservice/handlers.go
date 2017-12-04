package webservice

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// DEBUG es una constante que indica que se activa el modo depuración
// a lo largo de todo el codigo hay diferentes mensajes de pantalla
// para poder realizar un rustico depurado de la aplicación.
const DEBUG bool = false

// Esta es la dirección URL donde tenemos alojado nuestro servidor de NextCloud.
//
// Para realizar el despliegue automatico tenemos que tener en cuenta que debe
// haber una variable de entorno que se llame URL_BASE y que indique cual es
// la URL a donde se realizaran las peticiones para NextCloud.
var URL_BASE = os.Getenv("URL_BASE")

// Esta variable se usa para saber cual será el metodo por defecto que
// se utilizará para realizar las peticiones a NextCloud.
//
// Entre las opciones que tenemos destacan:
//
// PROPFIND, GET
//
// Las demas opciones las podemos ver en la documentación de NextCloud.
// https://docs.nextcloud.com/server/12/developer_manual/client_apis/WebDAV/index.html
var METHOD = "PROPFIND"

// Como estamos desarrollando una API que se encaga de comunicarnos con
// NextCloud, y la propia API de NextCloud se accede a través de un cliente
// web necesitamos generar un cliente en nuestra API para poder realizar
// las peticiones.
//
// Por eso con esta variable creamos un cliente el cual nos permite realizar
// peticiones sobre paginas https pero sin la necesidad de la verificaión
// del certificado de la pagina, esto se hace por si la pagina tiene un
// certificado auto-firmado.
var NET_CLIENT = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

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

// FList funcion que se encarga de
func FList(w http.ResponseWriter, r *http.Request) {
	if DEBUG {
		fmt.Println("Hacemos la peticion de la lista")
		fmt.Println("La URL es: " + URL_BASE)
	}
	peticion, err := http.NewRequest(METHOD, URL_BASE, nil)
	if err != nil {
		if DEBUG {
			fmt.Println("Este es el error de no encuentro la URL")
		}
		panic(err)
	}
	peticion.SetBasicAuth(os.Getenv("USER_NEXTCLOUD"), os.Getenv("PASS_NEXTCLOUD"))
	peticion.Header.Set("Cache-Control", "no-cache")
	peticion.Header.Set("Content-Type", "application/xml")
	peticion.Header.Set("Charset", "utf-8")

	respuesta, err := NET_CLIENT.Do(peticion)
	if err != nil {
		if DEBUG {
			fmt.Println("Este es el error de que no tiene usuario ni pass")
		}
		panic(err)
	}

	if respuesta.StatusCode == http.StatusMultiStatus {
		bodyRaw, err := ioutil.ReadAll(respuesta.Body)
		//bodyString := string(bodyRaw)
		//fmt.Println("Resultado de la peticion a la WEB\n" + bodyString)
		if err != nil {
			if DEBUG {
				fmt.Println("Error en pasar la respuesta a string")
			}
			panic(err)
		}

		// Esta es la estructura de XML que vamos a usar para almacenar
		// temporalmente los datos
		var xmlData Multistatus
		// Esta es la estructura de datos que vamos a enviar en formato
		// JSON al cliente, es una colección de objetos de JSON
		var responseJSON []JSONObject
		// Este es un evento de calandario en formato JSON
		var eventJSON JSONObject

		// en esta linea montamos un xml para luego poder navegar por el
		if err := xml.Unmarshal(bodyRaw, &xmlData); err != nil {
			fmt.Println("Este error es del XML decodificando")
			panic(err)
		}
		// con esto recorremos todo el XML
		for _, respTag := range xmlData.Responses {
			eventJSON = JSONObject{
				Id:       respTag.Propstat[0].Prop.Etag,
				Href:     respTag.Href,
				Name:     respTag.Propstat[0].Prop.Name,
				Modified: respTag.Propstat[0].Prop.Modified,
			}
			responseJSON = append(responseJSON, eventJSON)
			// con esto imprimimos por consola el resultado de XML
			if DEBUG {
				fmt.Println("URL: ", respTag.Href)
				for _, propstatTag := range respTag.Propstat {
					fmt.Println("\tStatus: ", propstatTag.Status)
					fmt.Println("\tName: ", propstatTag.Prop.Name)
					fmt.Println("\tContent-Type: ", propstatTag.Prop.Content_Type)
					fmt.Println("\tSize: ", propstatTag.Prop.Size)
					fmt.Println("\tModified: ", propstatTag.Prop.Modified)
					fmt.Println("\tEtag: ", propstatTag.Prop.Etag)
					for _, propTag := range propstatTag.Prop.PropList {
						fmt.Println("\t\tXMLName:", propTag.XMLName, "->", propTag.Value)
					}
					fmt.Println()
				}
				fmt.Println()
			}
		}
		// Con esto rellenamos el JSON que vamos a devolver al cliente
		if err := json.NewEncoder(w).Encode(responseJSON); err != nil {
			if DEBUG {
				fmt.Println("Este error es del JSON")
			}
			panic(err)
		}
	} else {
		if DEBUG {
			fmt.Println("El codigo de estado para la respuesta es:", respuesta.StatusCode)
		}
	}
	defer respuesta.Body.Close()
}

// se llama cuando se accede a "/show/{id}"
func FShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]
	json.NewEncoder(w).Encode(Id)
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

// se llama cuando se accede a "/exit"
func FExit(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
