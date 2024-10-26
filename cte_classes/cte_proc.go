package cte_classes

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"

	"CTe.Classes/Protocolo"
	"CTe.Classes/Servicos/Tipos"
)

// CteProc representa a estrutura do documento CTe.
type CteProc struct {
	XMLName xml.Name          `xml:"cteProc"`
	Versao  Tipos.Versao      `xml:"versao,attr"`
	CTe     CTe               `xml:"CTe"`
	ProtCTe Protocolo.ProtCTe `xml:"protCTe"`
}

// LoadXmlString carrega um objeto CteProc a partir de uma string XML.
func LoadXmlString(xmlStr string) (*CteProc, error) {
	var cteProc CteProc
	if err := xml.Unmarshal([]byte(xmlStr), &cteProc); err != nil {
		return nil, err
	}
	return &cteProc, nil
}

// LoadXmlFile carrega um objeto CteProc a partir de um arquivo XML.
func LoadXmlFile(filePath string) (*CteProc, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	return LoadXmlString(string(xmlData))
}
