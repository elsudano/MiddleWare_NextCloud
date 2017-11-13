package main

import (
	"encoding/json"
    "os"
	"net/http"
	"crypto/tls"

	"github.com/gorilla/mux"
)
// Variables de configuración
var URL_BASE = "https://www.sudano.net/nextcloud/remote.php/dav/calendars/prueba/personal/"
var METHOD = "PROPFIND" // PROPFIND, GET
var NET_CLIENT = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
}

// se llama cuando se accede a "/status" o a "/"
func FStatus(w http.ResponseWriter, r *http.Request) {
    status := Status{Status: "OK"}
    json.NewEncoder(w).Encode(status)
}

// se llama cuando se accede a "/list"
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
