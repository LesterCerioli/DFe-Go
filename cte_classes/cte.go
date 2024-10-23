package cte_classes

import (
	"encoding/xml"
	"errors"
	"os"

	"CTe.Classes/Ext"
	"CTe.Classes/Informacoes"
	"CTe.Classes/Servicos/Tipos"
	"DFe.Classes/Assinatura"
	"DFe.Utils"
)

// CTe representa a estrutura do CTe.
type CTe struct {
	XMLName      xml.Name                `xml:"CTe" xmlns:"http://www.portalfiscal.inf.br/cte"`
	Versao       *Tipos.Versao           `xml:"-"`
	ProxyVersao  string                  `xml:"versao,attr"`
	VersaoSpecified bool                 `xml:"-"`
	InfCte       Informacoes.InfCte      `xml:"infCte"`
	InfCTeSupl   Informacoes.InfCTeSupl  `xml:"infCTeSupl"`
	Signature    DFe.Assinatura.Signature `xml:"Signature"`
}

// GetString retorna a representação em string da versão
func (v *Tipos.Versao) GetString() string {
	switch *v {
	case Tipos.VersaoVe200:
		return "2.00"
	case Tipos.VersaoVe300:
		return "3.00"
	case Tipos.VersaoVe400:
		return "4.00"
	}
	return ""
}

// SetProxyVersao configura a versão com base na string
func (c *CTe) SetProxyVersao(value string) {
	switch value {
	case "2.00":
		c.Versao = &Tipos.VersaoVe200
	case "3.00":
		c.Versao = &Tipos.VersaoVe300
	case "4.00":
		c.Versao = &Tipos.VersaoVe400
	}
}

// LoadXmlString carrega um objeto CTe a partir de uma string XML.
func LoadXmlString(xmlStr string) (*CTe, error) {
	var cte CTe
	if err := xml.Unmarshal([]byte(xmlStr), &cte); err != nil {
		return nil, err
	}
	return &cte, nil
}

// LoadXmlFile carrega um objeto CTe a partir de um arquivo XML.
func LoadXmlFile(filePath string) (*CTe, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	xmlData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return LoadXmlString(string(xmlData))
}
