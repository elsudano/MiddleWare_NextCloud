package webservice

// Status este es un tipo de objeto especifico que se utiliza para
// realizar un test de disponibilidad del webservice.
type Status struct {
	Status string `json:"status"`
}

// JSONObject es un tipo que se encarga de almacenar los datos mas importantes
// de la respuesta que nos da NextCloud.
type EventJSON struct {
	Id       string `json:"id"`
	Etag     string `json:"etag"`
	Modified string `json:"last_modified"` // formatear la fecha para que sea YYYY/MM/DD HH:MM
}

// Reminder es un tipo que se encarga de almacenar los datos de los Recordatorios
// que pertenecen a un evento, para luego poder ofrecer alertas al usuario
type Reminder struct {
	Id           string     `json:"id"`
	Denomination string     `json:"denomination"`
	DateStart    string     `json:"datestart"`
	TimeStart    string     `json:"timestart"`
}

// icsJSON es un tipo que se encarga de almacenar los datos del evento que se
// envia al cliente, y los que se obtienen por parte del mismo, por que el ID
// del mensaje que se recibe por parte del cliente es vacio.
type icsJSON struct {
	Id           string     `json:"id"`
	Denomination string     `json:"denomination"`
	Description  string     `json:"description"`
	Reminders    []Reminder `json:"reminders"`
	DateStart    string     `json:"datestart"`
	DateEnd      string     `json:"dateend"`
	TimeStart    string     `json:"timestart"`
	TimeEnd      string     `json:"timeend"`
}
