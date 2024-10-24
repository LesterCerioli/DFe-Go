package event

import (
	"encoding/xml"
	"cte/tipos"
)

// DetEvento é a struct que representa os detalhes do evento.
type DetEvento struct {
	VersaoEvento    tipos.Versao   `xml:"versao,attr"`       // Atributo de versão do evento
	EventoContainer EventoContainer `xml:",innerxml"` // Container para eventos específicos
}

// EventoContainer é a interface para os diferentes tipos de eventos.
type EventoContainer interface {
	// Esse método pode ser implementado para serializar o evento corretamente
	EventoTipo() string
}

// EvCancCTe representa o evento de cancelamento do CTe.
type EvCancCTe struct {
	XMLName xml.Name `xml:"evCancCTe"`
	// Outros campos específicos do cancelamento
}

// EventoTipo retorna o tipo de evento.
func (e EvCancCTe) EventoTipo() string {
	return "evCancCTe"
}

// EvCCeCTe representa o evento de carta de correção do CTe.
type EvCCeCTe struct {
	XMLName xml.Name `xml:"evCCeCTe"`
	// Outros campos específicos da carta de correção
}

// EventoTipo retorna o tipo de evento.
func (e EvCCeCTe) EventoTipo() string {
	return "evCCeCTe"
}

// EvPrestDesacordo representa o evento de prestação em desacordo.
type EvPrestDesacordo struct {
	XMLName xml.Name `xml:"evPrestDesacordo"`
	// Outros campos específicos do desacordo
}

// EventoTipo retorna o tipo de evento.
func (e EvPrestDesacordo) EventoTipo() string {
	return "evPrestDesacordo"
}
