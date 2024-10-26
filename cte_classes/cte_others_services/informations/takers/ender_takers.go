package takers

import (
	"fmt"
	"strconv"
)

// EnderTomaOs representa a estrutura equivalente em Go da classe C# enderTomaOs.
type EnderTomaOs struct {
	XLgr     string   `json:"xLgr"`     // Logradouro
	Nro      string   `json:"nro"`      // Número
	XCpl     string   `json:"xCpl"`     // Complemento
	XBairro  string   `json:"xBairro"`  // Bairro
	CMun     int64    `json:"cMun"`     // Código do Município
	XMun     string   `json:"xMun"`     // Nome do Município
	CEP      *int64   `json:"CEP,omitempty"` // Código de Endereçamento Postal (opcional)
	UF       state   `json:"UF"`       // Unidade Federativa
	CPais    *int     `json:"cPais,omitempty"` // Código do País (opcional)
	XPais    string   `json:"xPais"`    // Nome do País
}

// ProxyCEP retorna o CEP formatado como uma string com 8 dígitos.
func (e *EnderTomaOs) ProxyCEP() string {
	if e.CEP != nil {
		return fmt.Sprintf("%08d", *e.CEP)
	}
	return ""
}

// SetCEP define o valor do CEP a partir de uma string.
func (e *EnderTomaOs) SetCEP(value string) error {
	cep, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	e.CEP = &cep
}

// ProxyUF retorna a sigla da UF.
func (e *EnderTomaOs) ProxyUF() string {
	return e.UF.GetSiglaUfString()
}

// SetUF define a UF a partir de uma string.
func (e *EnderTomaOs) SetUF(value string) {
	e.UF = UF.SiglaParaEstado(value)
}

// CPaisSpecified retorna se o campo cPais está especificado.
func (e *EnderTomaOs) CPaisSpecified() bool {
	return e.CPais != nil
}
