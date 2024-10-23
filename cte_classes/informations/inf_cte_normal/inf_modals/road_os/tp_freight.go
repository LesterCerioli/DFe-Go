package road_os

import (
	"encoding/xml"
)

// TpFretamento representa os tipos de fretamento.
type TpFretamento int

const (
	// Eventual representa o fretamento eventual.
	Eventual TpFretamento = 1
	// Continuo representa o fretamento contínuo.
	Continuo TpFretamento = 2
)

// MarshalXML customiza a serialização XML da enumeração TpFretamento.
func (t TpFretamento) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

// UnmarshalXML customiza a deserialização XML da enumeração TpFretamento.
func (t *TpFretamento) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := d.DecodeElement(&value, &start); err != nil {
		return err
	}
	switch value {
	case "1":
		*t = Eventual
	case "2":
		*t = Continuo
	default:
		return nil // ou retornar um erro, se preferir
	}
	return nil
}

// String retorna a representação em string da enumeração TpFretamento.
func (t TpFretamento) String() string {
	switch t {
	case Eventual:
		return "1"
	case Continuo:
		return "2"
	default:
		return "0"
	}
}
