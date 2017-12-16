package webservice

import (
	"encoding/xml"
)

// Dato dentro de la lista del XML del tipo xml.Name
type PropValue struct {
	XMLName xml.Name `xml:""`
	Value   string   `xml:",chardata"`
}

// Dato dentro de la lista del XML del tipo resourcetype
type ResoValue struct {
	Collection string `xml:"DAV: collection"`
	Calendar   string `xml:"urn:ietf:params:xml:ns:caldav calendar"`
}

// Dato dentro de la lista del XML del tipo urn:ietf:params:xml:ns:caldav
type CalValue struct {
	CalName  string `xml:""`
	CompName string `xml:"urn:ietf:params:xml:ns:caldav comp name"`
	Opaque   string `xml:"urn:ietf:params:xml:ns:caldav opaque"`
}

// Dato que se encarga de almacenar los datos del evento en cuestión
// siendo las otras estructuras para almacenar los metadatos y algunos
// datos extras.
type ICal struct {
}

// Tercer nivel de etiqueta dentro de la lista del XML
//
// Nota: se puede poner omitempty en los diferentes campos
// para que solo guarde aquellos que tienen contenido.
type Prop struct {
	Name         string      `xml:"DAV: displayname,omitempty"`
	Content_Type string      `xml:"DAV: getcontenttype,omitempty"`
	Size         string      `xml:"DAV: getcontentlength,omitempty"`
	Position     int         `xml:"http://apple.com/ns/ical/ calendar-order,omitempty"`
	Color        string      `xml:"http://apple.com/ns/ical/ calendar-color,omitempty"`
	Modified     string      `xml:"DAV: getlastmodified,omitempty"`
	Etag         string      `xml:"DAV: getetag,omitempty"`
	NEvents      int         `xml:"http://sabredav.org/ns sync-token,omitempty"`
	Owner        string      `xml:"http://owncloud.org/ns owner-principal,omitempty"`
	OwnerDN      string      `xml:"http://nextcloud.com/ns owner-displayname,omitempty"`
	ResoList     []ResoValue `xml:"DAV: resourcetype"`
	SupCalComSet []CalValue  `xml:"urn:ietf:params:xml:ns:caldav supported-calendar-component-set"`
	SheduCalTran []CalValue  `xml:"urn:ietf:params:xml:ns:caldav schedule-calendar-transp"`
	PropList     []PropValue `xml:",any"`
}

// Segundo nivel de etiqueta dentro de la lista del XML
type Propstat struct {
	Prop   *Prop  `xml:"DAV: prop"`
	Status string `xml:"DAV: status"`
}

// Primer nivel de etiqueta dentro de la lista del XML
type Response struct {
	Href     string     `xml:"DAV: href"`
	Propstat []Propstat `xml:"DAV: propstat"`
}

// Estructura principal para la lista del XML
type Multistatus struct {
	Responses []Response `xml:"DAV: response"`
}
