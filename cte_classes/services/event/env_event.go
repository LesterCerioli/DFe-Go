package event

import "encoding/xml"

// EnvEvento representa o XML raiz do envio de eventos
type EnvEvento struct {
	XMLName xml.Name      `xml:"envEvento"`
	Versao  string        `xml:"versao,attr"`   // Versão do leiaute
	IDLote  int           `xml:"idLote"`        // Identificador de controle do Lote de envio do Evento
	Eventos []EventoCTe   `xml:"evento"`        // Lista de eventos (um lote pode conter até 20 eventos)
}

// EventoCTe representa o evento específico (definição futura)
type EventoCTe struct {
	// Defina os campos de eventoCTe aqui
}
