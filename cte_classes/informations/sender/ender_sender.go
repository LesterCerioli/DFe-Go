package sender

import (
	"fmt"
	"strconv"
)

// EnderReme representa o endereço do remetente.
type EnderReme struct {
	XLgr   string  // Logradouro
	Nro    string  // Número
	XCpl   string  // Complemento
	XBairro string // Bairro
	CMun   int64   // Código do Município
	XMun   string  // Nome do Município
	CEP    int64   // Código do CEP
	UF     Estado  // Unidade Federativa
	CPais  *int    // Código do País (pode ser nulo)
	XPais  string  // Nome do País
}

// ProxyCEP retorna o CEP com zeros à esquerda, no formato de 8 dígitos.
func (e *EnderReme) ProxyCEP() string {
	return fmt.Sprintf("%08d", e.CEP)
}

// SetProxyCEP define o valor do CEP a partir de uma string formatada.
func (e *EnderReme) SetProxyCEP(value string) error {
	cep, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	e.CEP = cep
	return nil
}

// ProxyUF retorna a sigla do estado.
func (e *EnderReme) ProxyUF() string {
	return e.UF.GetSiglaUfString()
}

// SetProxyUF define o valor do Estado a partir da sigla.
func (e *EnderReme) SetProxyUF(value string) error {
	uf, err := SiglaParaEstado(value)
	if err != nil {
		return err
	}
	e.UF = uf
	return nil
}

// CPaisSpecified retorna verdadeiro se o campo CPais tiver um valor.
func (e *EnderReme) CPaisSpecified() bool {
	return e.CPais != nil
}
