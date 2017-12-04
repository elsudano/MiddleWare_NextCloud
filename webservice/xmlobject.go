package webservice

import (
	"encoding/xml"
)

type PropValue struct {
	XMLName xml.Name `xml:""`
	Value   string   `xml:",chardata"`
}

type Prop struct {
	PropList []PropValue `xml:",any"`
	//PropList     []PropValue `xml:"DAV: resourcetype,omitempty"`
    Name         string      `xml:"DAV: displayname"`
    Content_Type string      `xml:"DAV: getcontenttype"`
    Size         string      `xml:"DAV: getcontentlength"`
	Modified     string      `xml:"DAV: getlastmodified"`
    Etag         string      `xml:"DAV: getetag"`
}

type Propstat struct {
	Prop   *Prop   `xml:"DAV: prop"`
    Status string  `xml:"DAV: status"`
}

type Response struct {
	Href     string     `xml:"DAV: href"`
	Propstat []Propstat `xml:"DAV: propstat"`
}

type Multistatus struct {
	Responses []Response `xml:"DAV: response"`
}
