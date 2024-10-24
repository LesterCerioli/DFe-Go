package road_os

import (
	"encoding/xml"
)

// TpPropProp é um tipo definido para representar o tipo de proprietário.
type TpPropProp int

// Prop representa as informações do proprietário.
type Prop struct {
	CPF            string      `json:"CPF" xml:"CPF"`
	CNPJ           string      `json:"CNPJ" xml:"CNPJ"`
	RNTRC          string      `json:"RNTRC" xml:"RNTRC"`
	TAF            string      `json:"TAF" xml:"TAF"`
	NroRegEstadual string      `json:"NroRegEstadual" xml:"NroRegEstadual"`
	XNome          string      `json:"xNome" xml:"xNome"`
	IE             string      `json:"IE" xml:"IE"`
	UF             Estado      `json:"-" xml:"-"`
	TpProp        TpPropProp `json:"tpProp" xml:"tpProp"`
}

// ProxyUF retorna a sigla do estado como string.
func (p *Prop) ProxyUF() string {
	return p.UF.GetSiglaUfString()
}

// SetProxyUF define o estado a partir da sigla.
func (p *Prop) SetProxyUF(value string) {
	p.UF = UF.SiglaParaEstado(value)
}
