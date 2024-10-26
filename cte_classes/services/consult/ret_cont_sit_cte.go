package consult

import (
	"encoding/xml"
	"errors"
	"fmt"
)

// RetornoBase representa a estrutura base de retorno.
type RetornoBase struct {
	RetornoXmlString string
	EnvioXmlString   string
}

// retConsSitCTe representa a estrutura de retorno da consulta de situação do CTe.
type retConsSitCTe struct {
	XMLName       xml.Name        `xml:"retConsSitCTe" xmlns:"http://www.portalfiscal.inf.br/cte"`
	Versao        versao          `xml:"versao,attr"`
	TpAmb         TipoAmbiente    `xml:"tpAmb"`
	VerAplic      string          `xml:"verAplic"`
	CStat         int             `xml:"cStat"`
	XMotivo       string          `xml:"xMotivo"`
	CUF           Estado          `xml:"cUF"`
	ProtCTe       protCTe         `xml:"protCTe"`
	ProcEventoCTe []procEventoCTe `xml:"procEventoCTe"`
	RetornoBase                   // Embedding RetornoBase
}

// LoadXml carrega um XML e retorna um retConsSitCTe.
func LoadXml(xmlStr string) (*retConsSitCTe, error) {
	var retorno retConsSitCTe
	err := xml.Unmarshal([]byte(xmlStr), &retorno)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %w", err)
	}
	retorno.RetornoXmlString = xmlStr
	return &retorno, nil
}

// LoadXmlWithRequest carrega um XML e um objeto consSitCTe, retornando um retConsSitCTe.
func LoadXmlWithRequest(xmlStr string, consSitCTe consSitCTe) (*retConsSitCTe, error) {
	retorno, err := LoadXml(xmlStr)
	if err != nil {
		return nil, err
	}
	envioXmlStr, err := FuncoesXml.ClasseParaXmlString(consSitCTe)
	if err != nil {
		return nil, fmt.Errorf("error converting consSitCTe to XML: %w", err)
	}
	retorno.EnvioXmlString = envioXmlStr
	return retorno, nil
}
