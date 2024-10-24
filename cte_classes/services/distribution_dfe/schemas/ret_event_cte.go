package schemas

import "encoding/xml"

// RetEventoCTe representa a mensagem de retorno do resultado da solicitação de registro de evento do CT-e.
type RetEventoCTe struct {
	XMLName  xml.Name    `xml:"retEventoCTe" json:"retEventoCTe"`
	InfEvento RetInfEvento `xml:"infEvento" json:"infEvento"`
	Versao   float64     `xml:"versao,attr" json:"versao"`
}

// RetInfEvento representa a informação do evento dentro do retorno.
type RetInfEvento struct {
	// Defina os campos de acordo com a estrutura RetInfEvento
}
