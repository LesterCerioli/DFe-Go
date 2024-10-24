package reception

import (
	"encoding/xml"
)

// RetCTe representa a resposta do CTe
type RetCTe struct {
	// AR02 - Versão do leiaute
	Versao string `xml:"versao,attr"` // Atributo

	// AR03 - Identificação do Ambiente: 1 – Produção / 2 - Homologação
	TpAmb TipoAmbiente `xml:"tpAmb"`

	// AR06a - Código da UF que atendeu a solicitação
	CUF Estado `xml:"cUF"`

	// AR04 - Versão do Aplicativo que recebeu o Lote
	VerAplic string `xml:"verAplic"`

	// AR05 - Código do status da resposta
	CStat int `xml:"cStat"`

	// AR06 - Descrição literal do status da resposta
	DMotivo string `xml:"xMotivo"`

	// protCTe - Protocolo do CTe
	ProtCTe protCTe `xml:"protCTe"` // Certifique-se de que protCTe esteja definido
}

// LoadXml desserializa uma string XML para um objeto RetCTe
func LoadXml(xmlData string) (*RetCTe, error) {
	var retorno RetCTe
	err := xml.Unmarshal([]byte(xmlData), &retorno)
	if err != nil {
		return nil, err
	}
	return &retorno, nil
}

// LoadXmlWithEnvio desserializa uma string XML e associa com o objeto enviCTe
func LoadXmlWithEnvio(xmlData string, enviCte *CTe) (*RetCTe, error) {
	retorno, err := LoadXml(xmlData)
	if err != nil {
		return nil, err
	}
	// Implementar a lógica para associar enviCte, se necessário
	return retorno, nil
}
