package dispatcher

import (
	"encoding/xml"
	"fmt"
)

// Estado representa o estado (UF) com a lógica necessária.
type Estado struct {
	// Adicione campos e métodos conforme necessário
}

// GetSiglaUfString retorna a sigla do estado como string.
func (e Estado) GetSiglaUfString() string {
	// Implementar lógica para retornar a sigla do estado
	return "" // Substitua por lógica real
}

// SiglaParaEstado converte uma sigla em um objeto Estado.
func (e *Estado) SiglaParaEstado(sigla string) {
	// Implementar lógica para converter sigla em Estado
}

// enderExped representa o endereço do expedidor.
type enderExped struct {
	XLgr     string `xml:"xLgr"`
	Nro      string `xml:"nro"`
	XCpl     string `xml:"xCpl"`
	XBairro  string `xml:"xBairro"`
	CMun     int64  `xml:"cMun"`
	XMun     string `xml:"xMun"`
	CEP      int64  `xml:"-"`
	UF       Estado `xml:"-"`
	CPais    *int   `xml:"cPais,omitempty"` // Campo opcional, usa ponteiro para nil
	XPais    string `xml:"xPais"`
	ProxyCEP string `xml:"CEP"`
	ProxyUF  string `xml:"UF"`
}

// GetProxyCEP retorna o CEP formatado com zeros à esquerda.
func (e *enderExped) GetProxyCEP() string {
	return fmt.Sprintf("%08d", e.CEP)
}

// SetProxyCEP define o CEP a partir de uma string.
func (e *enderExped) SetProxyCEP(value string) error {
	var err error
	e.CEP, err = parseLong(value)
	return err
}

// GetProxyUF retorna a sigla do estado.
func (e *enderExped) GetProxyUF() string {
	return e.UF.GetSiglaUfString()
}

// SetProxyUF define o estado a partir de uma sigla.
func (e *enderExped) SetProxyUF(value string) {
	e.UF.SiglaParaEstado(value)
}

// parseLong converte uma string para int64, tratando erros.
func parseLong(value string) (int64, error) {
	var result int64
	_, err := fmt.Sscanf(value, "%d", &result)
	return result, err
}
