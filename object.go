package MiddleWare_NextCloud

// Status este es un tipo de objeto especifico que se utiliza para
// realizar un test de disponibilidad del webservice.
type Status struct {
    Status      string `json:"status"`
}

// Object es un tipo que se encarga de almacenar los datos mas importantes
// de la respuesta que nos da NextCloud.
type Object struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Completed  bool   `json:"completed"`
}

// Objects colección de Object en donde se almacenan los datos recibidos
// desde NextCloud.
type Objects []Object
