package stats

import (
	"encoding/xml"
	"time"
)

// RetornoBase deve ser definido com os campos e métodos necessários
type RetornoBase struct {
	RetornoXmlString   string // Substitua pelo tipo correto
	EnvioXmlString     string // Substitua pelo tipo correto
}

// RetConsStatServCte representa a resposta da consulta de status do CTe
type RetConsStatServCte struct {
	// FR02 - Versão do leiaute
	Versao string `xml:"versao,attr"` // Atributo

	// FR03 - Identificação do Ambiente: 1 – Produção / 2 - Homologação
	TpAmb TipoAmbiente `xml:"tpAmb"`

	// FR04 - Versão do Aplicativo que processou a consulta
	VerAplic string `xml:"verAplic"`

	// FR05 - Código do status da resposta
	CStat int `xml:"cStat"`

	// FR06 - Descrição literal do status da resposta
	XMotivo string `xml:"xMotivo"`

	// FR07 - Código da UF que atendeu a solicitação
	CUF Estado `xml:"cUF"`

	// Data e Hora do Recebimento
	DhRecbto time.Time `xml:"dhRecbto"` // Alterado para time.Time

	// Tempo médio
	TMed int `xml:"tMed"`

	// Data e Hora do Retorno
	DhRetorno time.Time `xml:"dhRetorno"` // Alterado para time.Time

	// Observação
	XObs string `xml:"xObs"`
}

// NewRetConsStatServCte cria uma nova instância de RetConsStatServCte
func NewRetConsStatServCte() *RetConsStatServCte {
	return &RetConsStatServCte{}
}

// LoadXml carrega a estrutura a partir de um XML
func (r *RetConsStatServCte) LoadXml(xmlStr string, consStatServCte ConsStatServCte) (*RetConsStatServCte, error) {
	retorno := &RetConsStatServCte{}
	err := xml.Unmarshal([]byte(xmlStr), retorno)
	if err != nil {
		return nil, err
	}
	retorno.RetornoXmlString = xmlStr
	retorno.EnvioXmlString = FuncoesXml.ClasseParaXmlString(consStatServCte)

	return retorno, nil
}

// RetConsStatServCTe representa a resposta da consulta de status do CTe
type RetConsStatServCTe struct {
	// FR02 - Versão do leiaute
	Versao string `xml:"versao,attr"` // Atributo

	// FR03 - Identificação do Ambiente: 1 – Produção / 2 - Homologação
	TpAmb TipoAmbiente `xml:"tpAmb"`

	// FR04 - Versão do Aplicativo que processou a consulta
	VerAplic string `xml:"verAplic"`

	// FR05 - Código do status da resposta
	CStat int `xml:"cStat"`

	// FR06 - Descrição literal do status da resposta
	XMotivo string `xml:"xMotivo"`

	// FR07 - Código da UF que atendeu a solicitação
	CUF Estado `xml:"cUF"`

	// Data e Hora do Recebimento
	DhRecbto time.Time `xml:"dhRecbto"` // Alterado para time.Time

	// Tempo médio
	TMed int `xml:"tMed"`

	// Data e Hora do Retorno
	DhRetorno time.Time `xml:"dhRetorno"` // Alterado para time.Time

	// Observação
	XObs string `xml:"xObs"`
}

// NewRetConsStatServCTe cria uma nova instância de RetConsStatServCTe
func NewRetConsStatServCTe() *RetConsStatServCTe {
	return &RetConsStatServCTe{}
}

// LoadXml carrega a estrutura a partir de um XML
func (r *RetConsStatServCTe) LoadXml(xmlStr string, consStatServCTe ConsStatServCTe) (*RetConsStatServCTe, error) {
	retorno := &RetConsStatServCTe{}
	err := xml.Unmarshal([]byte(xmlStr), retorno)
	if err != nil {
		return nil, err
	}
	retorno.RetornoXmlString = xmlStr
	retorno.EnvioXmlString = FuncoesXml.ClasseParaXmlString(consStatServCTe)

	return retorno, nil
}
