package MiddleWare_NextCloud

import (
	"encoding/json"
    "os"
	"net/http"
	"crypto/tls"

	"github.com/gorilla/mux"
)

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
    status := Status{Status: "OK"}
    json.NewEncoder(w).Encode(status)
}

// FList funcion que se encarga de
func FList(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(METHOD, URL_BASE, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(os.Getenv("USER_NEXTCLOUD"), os.Getenv("PASS_NEXTCLOUD"))
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Charset", "utf-8")

	resp, err := NET_CLIENT.Do(req)
	if err != nil {
		panic(err)
	}
	// objects := Objects{
	// 	Object{Name: "Write presentation"},
	// 	Object{Name: "Host meetup"},
	// }
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
	defer resp.Body.Close()
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
