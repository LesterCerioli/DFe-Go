package identification

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// enderToma representa o endereço do tomador.
type enderToma struct {
	XLgr     string `xml:"xLgr"`
	Nro      string `xml:"nro"`
	XCpl     string `xml:"xCpl"`
	XBairro  string `xml:"xBairro"`
	CMun     int64  `xml:"cMun"`
	XMun     string `xml:"xMun"`
	CEP      *int64 `xml:"-"`
	ProxyCEP string `xml:"CEP"`
	UF       Estado `xml:"-"`
	ProxyUF  string `xml:"UF"`
	CPais    *int   `xml:"cPais,omitempty"`
	XPais    string `xml:"xPais"`
}

// GetProxyCEP retorna o CEP formatado.
func (e *enderToma) GetProxyCEP() string {
	if e.CEP != nil {
		return fmt.Sprintf("%08d", *e.CEP)
	}
	return ""
}

// SetProxyCEP define o CEP a partir do valor string.
func (e *enderToma) SetProxyCEP(value string) error {
	if cep, err := strconv.ParseInt(value, 10, 64); err == nil {
		e.CEP = &cep
		return nil
	}
	return fmt.Errorf("invalid CEP value: %s", value)
}

// GetProxyUF retorna a sigla do estado.
func (e *enderToma) GetProxyUF() string {
	return e.UF.GetSiglaUfString()
}

// SetProxyUF define o estado a partir da sigla.
func (e *enderToma) SetProxyUF(value string) error {
	e.UF = UF.SiglaParaEstado(value)
	return nil
}

// CPaisSpecified verifica se cPais está especificado.
func (e *enderToma) CPaisSpecified() bool {
	return e.CPais != nil
}
