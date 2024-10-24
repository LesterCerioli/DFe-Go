package event

import (
	"encoding/xml"
	"cte/tipos"
	"dfes/assinatura"
)

// EventoCTe representa a estrutura do evento CTe
type EventoCTe struct {
	XMLName   xml.Name           `xml:"eventoCTe"`
	Versao    tipos.Versao       `xml:"versao,attr"`                        // Atributo versão do leiaute do evento
	InfEvento *InfEventoEnv      `xml:"infEvento"`                          // Grupo de informações do registro do Evento
	Signature *assinatura.Signature `xml:"Signature,omitempty"`             // Assinatura Digital do documento XML
}

// NewEventoCTe cria uma nova instância de EventoCTe
func NewEventoCTe(versao tipos.Versao, infEvento *InfEventoEnv, signature *assinatura.Signature) *EventoCTe {
	return &EventoCTe{
		Versao:    versao,
		InfEvento: infEvento,
		Signature: signature,
	}
}
