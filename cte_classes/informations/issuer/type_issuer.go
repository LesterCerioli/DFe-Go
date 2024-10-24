package issuer

import (
	"encoding/xml"
	"fmt"
)

// CRT representa a classificação do regime tributário.
type CRT int

const (
	// Simples Nacional
	SimplesNacional CRT = 1
	// Simples Nacional – excesso de sublimite de receita bruta
	SimplesNacionalExcessoSublimite CRT = 2
	// Regime Normal
	RegimeNormal CRT = 3
	// Simples Nacional MEI
	SimplesNacionalMei CRT = 4
)

// MarshalXML implementa a interface xml.Marshaler para CRT.
func (c CRT) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(c.String(), start)
}

// String retorna a representação em string do CRT.
func (c CRT) String() string {
	switch c {
	case SimplesNacional:
		return "Simples Nacional"
	case SimplesNacionalExcessoSublimite:
		return "Simples Nacional – excesso de sublimite de receita bruta"
	case RegimeNormal:
		return "Regime Normal"
	case SimplesNacionalMei:
		return "Simples Nacional MEI"
	default:
		return fmt.Sprintf("Unknown CRT value: %d", c)
	}
}

// UnmarshalXML implementa a interface xml.Unmarshaler para CRT.
func (c *CRT) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := d.DecodeElement(&value, &start); err != nil {
		return err
	}

	switch value {
	case "1":
		*c = SimplesNacional
	case "2":
		*c = SimplesNacionalExcessoSublimite
	case "3":
		*c = RegimeNormal
	case "4":
		*c = SimplesNacionalMei
	default:
		return fmt.Errorf("invalid CRT value: %s", value)
	}
	return nil
}
