package distribution_dfe

import "encoding/xml"

// ConsNSU representa o grupo para consultar um DF-e a partir de um NSU específico.
type ConsNSU struct {
	XMLName xml.Name `xml:"consNSU"`
	NSU     string   `xml:"NSU" json:"nsu"`
}
