package reception

import (
	"encoding/xml"
	"CTe/Classes/Servicos/Tipos"
	CteEletronica "CTe/Classes/CTe" // Ajuste o caminho de importação conforme necessário
)

// EnviCTe representa a estrutura de envio do CTe
type EnviCTe struct {
	Versao  Tipos.Versao              `xml:"versao,attr"` // AP02 - Versão do leiaute
	IDLote  int                       `xml:"idLote"`       // AP03 - Identificador de controle do envio do lote
	CTe     []CteEletronica.CTe      `xml:"CTe"`          // AP04 - Conjunto de CT-e transmitidas
}

// NewEnviCTe cria uma nova instância de EnviCTe
func NewEnviCTe(versao Tipos.Versao, idLote int, cTe []CteEletronica.CTe) *EnviCTe {
	return &EnviCTe{
		Versao: versao,
		IDLote: idLote,
		CTe:    cTe,
	}
}

// LoadXmlString carrega um EnviCTe a partir de uma string XML
func LoadXmlString(xmlStr string) (*EnviCTe, error) {
	var envi EnviCTe
	err := xml.Unmarshal([]byte(xmlStr), &envi)
	if err != nil {
		return nil, err
	}
	return &envi, nil
}

// LoadXmlArquivo carrega um EnviCTe a partir de um arquivo XML
func LoadXmlArquivo(caminhoArquivoXml string) (*EnviCTe, error) {
	xmlData, err := os.ReadFile(caminhoArquivoXml) // Use "os" para ler arquivos
	if err != nil {
		return nil, err
	}
	return LoadXmlString(string(xmlData))
}
