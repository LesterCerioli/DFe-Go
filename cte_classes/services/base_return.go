package services

import "encoding/xml"

// RetornoBase representa a estrutura base para retorno de serviços.
type RetornoBase struct {
	EnvioXmlString  string `xml:"-"`
	RetornoXmlString string `xml:"-"`
}
