package webservice

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/apognu/gocal"
)

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

// Estructura de datos que se encarga de mantener la lectura del XML para que
// nos sirva como base de datos temporal
var xmlData = readXML()

// readXML Con esta función leemos directamente desde una fuente WebDAV un calendario
// y lo comvertimos a una estructura XML que podemos leer despues desde otras
// funcionalidades para poder presentar la información de otras maneras.
//
// El formato en el que se devuelve este XML viene dado por una serie de
// estructuras que se han definido en el fichero xmlobject.go
func readXML() Multistatus {
	if DEBUG {
		fmt.Println("Hacemos la peticion de la lista")
		fmt.Println("La URL Completa es: " + URL_COMPLETE_PATH)
	}
	// Esta es la estructura de XML que vamos a usar para almacenar
	// temporalmente los datos
	var xmlData Multistatus
	// se almacena la petición que se realiza a la URL_COMPLETE_PATH y
	// un err por si falla la petición
	peticion, err := http.NewRequest("PROPFIND", URL_COMPLETE_PATH, nil)
	if err != nil {
		if DEBUG {
			fmt.Println("Este es el error de no encuentro la URL")
		}
		panic(err)
	}
	// como no da ningun error se procede a setear las diferentes cabeceras
	// y demas parametros de la petición
	peticion.SetBasicAuth(os.Getenv("USER_NEXTCLOUD"), os.Getenv("PASS_NEXTCLOUD"))
	peticion.Header.Set("Cache-Control", "no-cache")
	peticion.Header.Set("Content-Type", "application/xml")
	peticion.Header.Set("Charset", "utf-8")
	// al igual que la solicitud tambien se hace con la respuesta
	// se genera un err por si falla dicha respuesta
	respuesta, err := NET_CLIENT.Do(peticion)
	if err != nil {
		if DEBUG {
			fmt.Println("Este es el error de que no tiene usuario ni pass")
		}
		panic(err)
	}
	// normalmente en los WebDAV el estado devuelto suele ser MultiEstado
	// en http://sabre.io/dav/building-a-caldav-client/ se puede ver
	if respuesta.StatusCode == http.StatusMultiStatus {
		bodyRaw, err := ioutil.ReadAll(respuesta.Body)
		if err != nil {
			if DEBUG {
				fmt.Println("Error en pasar la respuesta a string")
			}
			panic(err)
		}
		// en esta linea montamos un xml para luego poder navegar por el
		if err := xml.Unmarshal(bodyRaw, &xmlData); err != nil {
			fmt.Println("Este error es del XML decodificando")
			panic(err)
		}
	} else if DEBUG {
		fmt.Println("El codigo de estado para la respuesta es:", respuesta.StatusCode)
	}
	defer respuesta.Body.Close()
	return xmlData
}

// readICS Con esta función leemos el fichero ICS desde una ruta de internet
// y lo devolvemos en un formato JSON
//
// Para que se pueda realizar cualquier operación con el fichero descargado
// se devuelve una cadena y de esa manera se puede parsear o guardar.
func readICS(icsFILE string) icsJSON {
	// esta es la información que devolvemos
	var JSONstruct icsJSON
	if DEBUG {
		fmt.Println("El evento es: " + icsFILE)
	}
	// se almacena la petición que se realiza a la URL_COMPLETE_PATH y
	// un err por si falla la petición
	peticion, err := http.NewRequest("GET", icsFILE, nil)
	if err != nil {
		if DEBUG {
			fmt.Println("Este es el error de no encuentro la URL")
		}
		panic(err)
	}
	// como no da ningun error se procede a setear las diferentes cabeceras
	// y demas parametros de la petición
	peticion.SetBasicAuth(os.Getenv("USER_NEXTCLOUD"), os.Getenv("PASS_NEXTCLOUD"))
	peticion.Header.Set("Cache-Control", "no-cache")
	peticion.Header.Set("Content-Type", "text/calendar; charset=utf-8; component=vevent")
	peticion.Header.Set("Charset", "utf-8")
	// al igual que la solicitud tambien se hace con la respuesta
	// se genera un err por si falla dicha respuesta
	respuesta, err := NET_CLIENT.Do(peticion)
	if err != nil {
		panic(err)
	}
	// datos en bruto del fichero descargado
	fileRaw, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		panic(err)
	}
	// Creamos el fichero temporal
	tmpfile, err := ioutil.TempFile("", "ics-tmp-json-")
	if err != nil {
		panic(err)
	}
	// Escribimos en el fichero temporal el que nos descargamos
	if _, err := tmpfile.Write(fileRaw); err != nil {
		panic(err)
	}
	// ponemos la posicion de lectura al principio del fichero
	tmpfile.Seek(0, 0)
	// ponemos el fichero en un calendario
	calendar := gocal.NewParser(tmpfile)
	calendar.Parse()
	// si el fichero es un evento solo tendrá 1 elemento
	if len(calendar.Events) == 1 {
		event := calendar.Events[0]
		JSONstruct = icsJSON {
			Id: event.Uid,
			Denomination: event.Summary,
			Description: event.Description,
			Start: event.Start.String(),
			End: event.End.String(),
		}
		if DEBUG {
			fmt.Println(calendar.Events[0].Uid)
			fmt.Println(calendar.Events[0].Summary)
			fmt.Println(calendar.Events[0].Description)
		}
	// si el fichero es un calendario tendrá varios elementos
	} else if len(calendar.Events) > 1 {
		for _, event := range calendar.Events {
			if DEBUG {
				fmt.Printf("%s on %s by %s", event.Summary, event.Start.String(), event.Organizer.Cn)
			}
		}
	}
	// Cerramos el fichero temporal para que se borre
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name()) // lo preparamos para ser borrado
	defer respuesta.Body.Close()
	return JSONstruct
}

// print
// printXML Con esta funcion podemos imprimir el contenido del XML que pasamos
// por parametros.
//
// Esta función solo imprime por pantalla el contenido del mismo.
func printXML(xmlData Multistatus) {
	for i, respTag := range xmlData.Responses {
		if i == 0 {
			// El response 0 es la posición donde están los datos del calendario
			fmt.Println("Esta es la información del calendario\nURL:", respTag.Href)
			for _, propstatTag := range respTag.Propstat {
				fmt.Println("\tStatus:", propstatTag.Status)
				fmt.Println("\tObject Display Name:", propstatTag.Prop.Name)
				fmt.Println("\tPosition:", propstatTag.Prop.Position)
				fmt.Println("\tColor:", propstatTag.Prop.Color)
				// le quitamos 1 por que el primero es el registro del calendario
				fmt.Println("\tNumero de Eventos:", propstatTag.Prop.NEvents-1)
				fmt.Println("\tOwner Principal:", propstatTag.Prop.Owner)
				fmt.Println("\tOwner Display Name:", propstatTag.Prop.OwnerDN)
				fmt.Println("\tResourceType List:")
				for _, ResoTag := range propstatTag.Prop.ResoList {
					fmt.Println("\t\tCollection:", ResoTag.Collection)
					fmt.Println("\t\tCalendar:", ResoTag.Calendar)
				}
				fmt.Println("\tSupport Calendar Component Set:", propstatTag.Prop.SupCalComSet)
				fmt.Println("\tShedule Calendar Transport:", propstatTag.Prop.SheduCalTran)
				fmt.Println("\tOthers Visible options:")
				for _, propTag := range propstatTag.Prop.PropList {
					fmt.Println("\t\tXMLName:", propTag.XMLName, "->", propTag.Value)
				}
				fmt.Println()
			}
			fmt.Println()
		} else if respTag.Propstat[0].Status == STATUSOK && strings.Contains(respTag.Propstat[0].Prop.Content_Type, "component=vevent") {
			// en esta parte mostramos el resultado del JSON y formateamos la
			// salida por consola de un objeto tipo evento
			fmt.Println("Esta es la información de un evento\nURL:", respTag.Href)
			for _, propstatTag := range respTag.Propstat {
				if propstatTag.Status == STATUSOK {
					fmt.Println("\tEsto es un prop con Status:", propstatTag.Status)
					fmt.Println("\tContent-Type:", propstatTag.Prop.Content_Type)
					fmt.Println("\tSize:", propstatTag.Prop.Size)
					fmt.Println("\tModified:", propstatTag.Prop.Modified)
					// le quitamos 1 por que el primero es el registro del calendario
					fmt.Println("\tEtag:", strings.Replace(propstatTag.Prop.Etag, "\"", "", -1))
					fmt.Println("\tResourceType List:")
					for _, ResoTag := range propstatTag.Prop.ResoList {
						fmt.Println("\t\tCollection:", ResoTag.Collection)
						fmt.Println("\t\tCalendar:", ResoTag.Calendar)
					}
				} else {
					fmt.Println("\tEsto es un prop con Status:", propstatTag.Status)
					fmt.Println("\tOthers Visible options:")
					for _, propTag := range propstatTag.Prop.PropList {
						fmt.Println("\t\tXMLName:", propTag.XMLName, "->", propTag.Value)
					}
				}

			}
			fmt.Println()
		} else {
			// en esta parte imprimimos todos los demas elementos
			// del XML que no hemos contemplado arriba
			fmt.Println("Esto es algo que no estaba contemplado\nURL:", respTag.Href)
			for _, propstatTag := range respTag.Propstat {
				fmt.Println("\tStatus:", propstatTag.Status)
				fmt.Println("\tObject Display Name:", propstatTag.Prop.Name)
				fmt.Println("\tContent-Type:", propstatTag.Prop.Content_Type)
				fmt.Println("\tSize:", propstatTag.Prop.Size)
				fmt.Println("\tPosition:", propstatTag.Prop.Position)
				fmt.Println("\tColor:", propstatTag.Prop.Color)
				fmt.Println("\tModified:", propstatTag.Prop.Modified)
				// le quitamos 1 por que el primero es el registro del calendario
				fmt.Println("\tNumero de Eventos:", propstatTag.Prop.NEvents-1)
				fmt.Println("\tEtag:", strings.Replace(propstatTag.Prop.Etag, "\"", "", -1))
				fmt.Println("\tOwner Principal:", propstatTag.Prop.Owner)
				fmt.Println("\tOwner Display Name:", propstatTag.Prop.OwnerDN)
				fmt.Println()
				fmt.Println("\tResourceType List:")
				for _, ResoTag := range propstatTag.Prop.ResoList {
					fmt.Println("\t\tCollection:", ResoTag.Collection)
					fmt.Println("\t\tCalendar:", ResoTag.Calendar)
				}
				fmt.Println("\tSupport Calendar Component Set:", propstatTag.Prop.SupCalComSet)
				fmt.Println("\tShedule Calendar Transport:", propstatTag.Prop.SheduCalTran)
				fmt.Println("\tOthers Visible options:")
				for _, propTag := range propstatTag.Prop.PropList {
					fmt.Println("\t\tXMLName:", propTag.XMLName, "->", propTag.Value)
				}
				fmt.Println()
			}
			fmt.Println()
		}
	}
}
