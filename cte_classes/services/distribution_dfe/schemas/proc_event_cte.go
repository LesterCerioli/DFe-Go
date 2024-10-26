package schemas

import "encoding/xml"

// ProcEventoCTe representa o leiaute de compartilhamento de solicitação de registro de evento do CT-e.
type ProcEventoCTe struct {
	XMLName       xml.Name    `xml:"procEventoCTe" json:"procEventoCTe"`
	Versao        float64     `xml:"versao,attr" json:"versao"`
	IPTransmissor string      `xml:"ipTransmissor,attr" json:"ipTransmissor"`
	EventoCTe     EventoCTe   `xml:"eventoCTe" json:"eventoCTe"`
	RetEventoCTe  RetEventoCTe `xml:"retEventoCTe" json:"retEventoCTe"`
}
