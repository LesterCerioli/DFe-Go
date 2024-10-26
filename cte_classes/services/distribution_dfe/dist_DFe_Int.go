package distribution_dfe

import (
	"encoding/xml"
	"errors"
)

// Constantes
const ErroCpfCnpjPreenchidos = "Somente preencher um dos campos: CNPJ ou CPF, para um objeto do tipo dest!"

// DistDFeInt representa a estrutura de distribuição de DF-e de interesse.
type DistDFeInt struct {
	XMLName   xml.Name    `xml:"distDFeInt"`
	Versao    string      `xml:"versao,attr"`
	TpAmb     TipoAmbiente `xml:"tpAmb"`
	CUFAutor  Estado      `xml:"cUFAutor"`
	CNPJ      string      `xml:"CNPJ,omitempty"`
	CPF       string      `xml:"CPF,omitempty"`
	DistNSU   *DistNSU    `xml:"distNSU,omitempty"`
	ConsNSU   *ConsNSU    `xml:"consNSU,omitempty"`
}

// SetCNPJ define o valor de CNPJ garantindo que CPF não esteja preenchido.
func (d *DistDFeInt) SetCNPJ(cnpj string) error {
	if cnpj != "" && d.CPF != "" {
		return errors.New(ErroCpfCnpjPreenchidos)
	}
	d.CNPJ = cnpj
	return nil
}

// SetCPF define o valor de CPF garantindo que CNPJ não esteja preenchido.
func (d *DistDFeInt) SetCPF(cpf string) error {
	if cpf != "" && d.CNPJ != "" {
		return errors.New(ErroCpfCnpjPreenchidos)
	}
	d.CPF = cpf
	return nil
}
