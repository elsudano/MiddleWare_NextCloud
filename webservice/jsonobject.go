package webservice

// Status este es un tipo de objeto especifico que se utiliza para
// realizar un test de disponibilidad del webservice.
type Status struct {
    Status      string `json:"status"`
}

// JSONObject es un tipo que se encarga de almacenar los datos mas importantes
// de la respuesta que nos da NextCloud.
type JSONObject struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
    Href       string `json:"href"`
    Modified   string `json:"last_modified"`
	Completed  bool   `json:"completed"`
}
