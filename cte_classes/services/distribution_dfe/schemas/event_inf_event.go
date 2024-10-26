package schemas

import (
	"encoding/xml"
	"time"
)

// EventoInfEvento representa as informações do evento no CTe.
type EventoInfEvento struct {
	XMLName     xml.Name   `xml:"infEvento" json:"infEvento"`
	Id          string     `xml:"Id,attr" json:"Id"`
	COrgao      byte       `xml:"cOrgao" json:"cOrgao"`
	TpAmb       byte       `xml:"tpAmb" json:"tpAmb"`
	CNPJ        string     `xml:"CNPJ" json:"CNPJ"`
	ChCTe       string     `xml:"chCTe" json:"chCTe"`
	DhRecbto    time.Time  `xml:"-" json:"dhRecbto"` // Ignora no XML
	ProxyDhRecbto string   `xml:"dhRecbto" json:"dhRecbto"` // Utilizado para o valor serializado
	TpEvento    uint       `xml:"tpEvento" json:"tpEvento"`
	NSeqEvento  byte       `xml:"nSeqEvento" json:"nSeqEvento"`
	VersaoEvento float64   `xml:"versaoEvento" json:"versaoEvento"`
	DetEvento   DetEvento  `xml:"detEvento" json:"detEvento"`
}

// FormatTimeUTC é responsável por formatar a data e hora no formato UTC.
func FormatTimeUTC(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05Z")
}

// ParseTimeUTC é responsável por converter uma string em formato UTC para o tipo time.Time.
func ParseTimeUTC(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}

// MarshalXML é usado para garantir que o campo ProxyDhRecbto tenha o valor formatado corretamente.
func (e *EventoInfEvento) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	e.ProxyDhRecbto = FormatTimeUTC(e.DhRecbto)
	type Alias EventoInfEvento
	return encoder.EncodeElement((*Alias)(e), start)
}

// UnmarshalXML é usado para garantir que o campo DhRecbto seja devidamente convertido ao desserializar.
func (e *EventoInfEvento) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type Alias EventoInfEvento
	aux := &struct {
		*Alias
		ProxyDhRecbto string `xml:"dhRecbto"`
	}{
		Alias: (*Alias)(e),
	}
	if err := decoder.DecodeElement(aux, &start); err != nil {
		return err
	}
	parsedTime, err := ParseTimeUTC(aux.ProxyDhRecbto)
	if err != nil {
		return err
	}
	e.DhRecbto = parsedTime
	return nil
}
