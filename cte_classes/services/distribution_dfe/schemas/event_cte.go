package schemas

import (
	"encoding/xml"
)

// EventoCTe representa o evento do CTe.
type EventoCTe struct {
	InfEvento EventoInfEvento `xml:"infEvento" json:"infEvento"`
	Versao    float64         `xml:"versao,attr" json:"versao"`
	Signature Signature       `xml:"Signature" json:"Signature"`
}

// EventoInfEvento representa as informações do evento.
type EventoInfEvento struct {
	// Defina as propriedades da struct de acordo com os campos da infEvento.
}

// Signature representa a assinatura digital do documento XML.
type Signature struct {
	// Defina as propriedades da struct Signature, se necessário, ou importe uma biblioteca que forneça a implementação.
}
