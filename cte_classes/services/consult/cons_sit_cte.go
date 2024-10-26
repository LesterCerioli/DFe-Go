package consult

import "encoding/xml"

// consSitCTe representa a estrutura da consulta de situação do CTe.
type consSitCTe struct {
	XMLName xml.Name           `xml:"consSitCTe" xmlns:"http://www.portalfiscal.inf.br/cte"`
	Versao  *Tipos.Versao      `xml:"versao,attr"`
	TpAmb   Tipos.TipoAmbiente `xml:"tpAmb"`
	XServ   string             `xml:"xServ"`
	ChCTe   string             `xml:"chCTe"`
}

// NewConsSitCTe cria uma nova instância de consSitCTe com valores padrão.
func NewConsSitCTe() *consSitCTe {
	return &consSitCTe{
		XServ: "CONSULTAR",
	}
}
