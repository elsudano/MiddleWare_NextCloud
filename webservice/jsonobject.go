package webservice

// Status este es un tipo de objeto especifico que se utiliza para
// realizar un test de disponibilidad del webservice.
type Status struct {
    Status      string `json:"status"`
}

// JSONObject es un tipo que se encarga de almacenar los datos mas importantes
// de la respuesta que nos da NextCloud.
type EventJSON struct {
	Id           string `json:"id"`
    Etag         string `json:"etag"`
    Modified     string `json:"last_modified"`
}

// icsJSON es un tipo que se encarga de almacenar los datos del evento.
type icsJSON struct {
	Id           string `json:"id"`
    Denomination string `json:"denomination"`
    Description  string `json:"description"`
}
