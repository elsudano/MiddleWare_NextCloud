package webservice

import (
	"os"
)

// Este es el protocolo que se usara para la comunicación con el servidor
//
// Las diferentes opciones que tenemos son: http o https
var PROTOCOL = "https"

// Esta variable indica cual sera el dominio de primer nivel en donde se
// encuentra el servidor de NextCloud
//
// Hay que tener en cuenta que si nuestro dominio tiene las "www" hay que
// añadirlo para que al componer la ruta completa donde se encuentra
// el servicio de NextCloud seamos capaces de localizarlo.
var DOMAIN = os.Getenv("DOMAIN")

// Esta es la dirección URL donde tenemos alojado nuestro servidor de NextCloud.
//
// Para realizar el despliegue automatico tenemos que tener en cuenta que debe
// haber una variable de entorno que se llame URL_BASE y que indique cual es
// la URL a donde se realizaran las peticiones para NextCloud.
var URL_BASE = os.Getenv("URL_BASE")

// Esta es la variable que almacena la ruta completa donde se encuentra
// el calendario a donde queremos acceder
//
// Se realiza la separación de todos los componentes de la dirección para
// despues utilizarlos en las diferentes partes del programa.
var URL_COMPLETE_PATH = PROTOCOL + "://" + DOMAIN + URL_BASE
