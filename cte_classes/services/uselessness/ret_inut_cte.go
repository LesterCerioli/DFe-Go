package uselessness

import (
	"encoding/xml"
	"DFe/Classes/Assinatura" // Ajuste o caminho de importação conforme necessário
	"CTe/Classes/Servicos/Tipos" // Ajuste o caminho de importação conforme necessário
	"DFe/Utils" // Ajuste o caminho de importação conforme necessário
)

// RetInutCTe representa a estrutura da resposta de inutilização do CTe
type RetInutCTe struct {
	Versao   Versao                     `xml:"versao,attr"` // DR02 - Versão do leiaute
	InfInut  InfInutRet                 `xml:"infInut"`     // DR03 - Dados da resposta
	Signature Assinatura.Signature       `xml:"Signature"`    // DR18 - Assinatura XML
}

// LoadXml carrega a estrutura RetInutCTe a partir de um XML e um objeto inutCTe
func LoadXml(xml string, inutCTe InutCTe) (*RetInutCTe, error) {
	retorno, err := loadXml(xml)
	if err != nil {
		return nil, err
	}
	// Converte inutCTe para XML e armazena em retorno.EnvioXmlString
	retorno.EnvioXmlString = Utils.ClasseParaXmlString(inutCTe) // Ajuste conforme a função de conversão

	return retorno, nil
}

// loadXml carrega a estrutura RetInutCTe a partir de um XML
func loadXml(xml string) (*RetInutCTe, error) {
	var retorno RetInutCTe
	err := xml.Unmarshal([]byte(xml), &retorno)
	if err != nil {
		return nil, err
	}
	retorno.RetornoXmlString = xml // Presumindo que você tenha esse campo

	return &retorno, nil
}
