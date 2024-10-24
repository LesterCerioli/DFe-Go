package stats

import (
	"encoding/xml"
)

// ConsStatServCte representa a solicitação de status do CTe
type ConsStatServCte struct {
	// FP02 - Versão do leiaute
	Versao string `xml:"versao,attr"` // Atributo

	// FP03 - Identificação do Ambiente: 1 – Produção / 2 - Homologação
	TpAmb TipoAmbiente `xml:"tpAmb"`

	// FP04 - Serviço solicitado 'STATUS'
	XServ string `xml:"xServ"`
}

// NewConsStatServCte cria uma nova instância de ConsStatServCte
func NewConsStatServCte() *ConsStatServCte {
	return &ConsStatServCte{
		XServ: "STATUS",
	}
}

// ConsStatServCTe representa a solicitação de status do CTe
type ConsStatServCTe struct {
	// FP02 - Versão do leiaute
	Versao string `xml:"versao,attr"` // Atributo

	// FP03 - Identificação do Ambiente: 1 – Produção / 2 - Homologação
	TpAmb TipoAmbiente `xml:"tpAmb"`

	// Código da UF que atendeu a solicitação
	CUF Estado `xml:"cUF"`

	// FP04 - Serviço solicitado 'STATUS'
	XServ string `xml:"xServ"`
}

// NewConsStatServCTe cria uma nova instância de ConsStatServCTe
func NewConsStatServCTe() *ConsStatServCTe {
	return &ConsStatServCTe{
		XServ: "STATUS",
	}
}
