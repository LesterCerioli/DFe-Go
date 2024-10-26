package consult

import "encoding/xml"

// procEventoCTe representa a estrutura do processamento do evento CTe.
type procEventoCTe struct {
	XMLName   xml.Name     `xml:"procEventoCTe" xmlns:"http://www.portalfiscal.inf.br/cte"`
	Versao    string       `xml:"versao,attr"`
	EventoCTe eventoCTe    `xml:"eventoCTe"`
	RetEvento retEventoCTe `xml:"retEventoCTe"`
}
