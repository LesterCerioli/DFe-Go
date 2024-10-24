package return

import (
	"encoding/xml"
	"CTe/Classes/Protocolo" // Ajuste o caminho de importação conforme necessário
	"CTe/Classes/Servicos/Tipos"
	"CTe/DFe/Classes/Entidades" // Ajuste o caminho de importação conforme necessário
	"CTe/DFe/Classes/Flags"     // Ajuste o caminho de importação conforme necessário
)

// RetConsReciCTe representa a estrutura da resposta da consulta do recibo do CTe
type RetConsReciCTe struct {
	Versao     Versao                   `xml:"versao,attr"` // BR02 - Versão do leiaute
	TpAmb      TipoAmbiente             `xml:"tpAmb"`       // Identificação do Ambiente
	VerAplic   string                   `xml:"verAplic"`    // Versão do Aplicativo que recebeu a Consulta
	NRec       string                   `xml:"nRec"`        // Número do Recibo consultado
	CStat      int                      `xml:"cStat"`       // Código do status da resposta
	CMotivo    string                   `xml:"xMotivo"`     // Descrição do status da resposta
	CUF        Estado                   `xml:"cUF"`         // Código da UF que atendeu a solicitação
	CMsg       *int                     `xml:"cMsg,omitempty"` // Código da Mensagem (v2.0)
	XMsg       string                   `xml:"xMsg"`        // Mensagem da SEFAZ para o emissor
	ProtCTe    []ProtCTe                `xml:"protCTe"`     // Conjunto de resultados do processamento
}

// LoadXml carrega um RetConsReciCTe a partir de uma string XML
func LoadXml(xmlStr string) (*RetConsReciCTe, error) {
	var retorno RetConsReciCTe
	err := xml.Unmarshal([]byte(xmlStr), &retorno)
	if err != nil {
		return nil, err
	}
	return &retorno, nil
}

// LoadXmlWithRequest carrega um RetConsReciCTe a partir de uma string XML e um consReciCTe
func LoadXmlWithRequest(xmlStr string, consSitCTe ConsReciCTe) (*RetConsReciCTe, error) {
	retorno, err := LoadXml(xmlStr)
	if err != nil {
		return nil, err
	}
	// Aqui você pode adicionar a lógica para serializar consSitCTe em XML se necessário
	// retorno.EnvioXmlString = FuncoesXml.ClasseParaXmlString(consSitCTe) // Exemplo
	return retorno, nil
}
