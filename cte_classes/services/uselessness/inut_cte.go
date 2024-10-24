package uselessness

import (
	"encoding/xml"
	"DFe/Classes/Assinatura" // Ajuste o caminho de importação conforme necessário
	"CTe/Classes/Servicos/Tipos" // Ajuste o caminho de importação conforme necessário
)

// InutCTe representa a estrutura de inutilização do CTe
type InutCTe struct {
	Versao   Versao                     `xml:"versao,attr"` // DP02 - Versão do leiaute
	InfInut  InfInutEnv                 `xml:"infInut"`     // DP03 - Dados do Pedido
	Signature Assinatura.Signature       `xml:"Signature"`    // DP15 - Assinatura XML
}

// NewInutCTe cria uma nova instância de InutCTe
func NewInutCTe() *InutCTe {
	return &InutCTe{
		InfInut: InfInutEnv{}, // Inicializa InfInutEnv
	}
}
