package receiver

import (
	"encoding/xml"
	"fmt"
)

// LocEnt representa a localização do destinatário.
type LocEnt struct {
	CNPJ    string `xml:"CNPJ"`    // CNPJ do destinatário
	CPF     string `xml:"CPF"`     // CPF do destinatário
	XNome   string `xml:"xNome"`   // Nome do destinatário
	XLgr    string `xml:"xLgr"`    // Logradouro
	Nro     string `xml:"nro"`     // Número
	XCpl    string `xml:"xCpl"`    // Complemento
	XBairro string `xml:"xBairro"` // Bairro
	CMun    int64  `xml:"cMun"`    // Código do Município
	XMun    string `xml:"xMun"`    // Nome do Município
	UF      Estado `xml:"-"`       // UF (não serializado diretamente)
	ProxyUF string `xml:"UF"`      // Proxy para a sigla da UF
}

// GetProxyUF retorna a sigla da UF.
func (l *LocEnt) GetProxyUF() string {
	return l.UF.GetSiglaUfString()
}

// SetProxyUF define a UF a partir de uma string.
func (l *LocEnt) SetProxyUF(value string) {
	l.UF = UF.SiglaParaEstado(value)
}

// Estado representa a sigla da UF (defina esta estrutura conforme necessário).
type Estado struct {
	// Defina os campos apropriados para Estado aqui
}

// GetSiglaUfString retorna a sigla da UF como string.
func (e Estado) GetSiglaUfString() string {
	// Implementar a lógica para obter a sigla da UF
	return "" // Retornar a sigla correspondente
}

// SiglaParaEstado converte uma sigla para um Estado.
func (e Estado) SiglaParaEstado(sigla string) Estado {
	// Implementar a lógica para converter a sigla para o Estado correspondente
	return Estado{} // Retornar o estado correspondente
}
