package broneAuth

import "encoding/xml"

type SAMLResponse struct {
	XMLName            xml.Name           `xml:"Response"`
	AttributeStatement AttributeStatement `xml:"Assertion>AttributeStatement"`
}

type AttributeStatement struct {
	Attributes []Attribute `xml:"Attribute"`
}

type Attribute struct {
	FriendlyName   string `xml:"FriendlyName,attr"`
	Name           string `xml:"Name,attr"`
	NameFormat     string `xml:"NameFormat,attr"`
	AttributeValue string `xml:"AttributeValue"`
}
