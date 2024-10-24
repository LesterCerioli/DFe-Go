package informations

import (
	"github.com/shopspring/decimal"
)

// InfCteOS representa a estrutura equivalente em Go da classe C# infCteOS.
type InfCteOS struct {
	Versao     VersaoServico    `json:"versao"`     // Versão do serviço
	ID         string            `json:"id"`         // Identificação
	Ide        IdeOs            `json:"ide"`        // Identificação do CTe
	Compl      ComplOs          `json:"compl"`      // Complementos
	Emit       EmitOs           `json:"emit"`       // Emitente
	Toma       TomaOs           `json:"toma"`       // Tomador
	VPrest     VPrestOs         `json:"vPrest"`     // Valores da Prestação
	Imp        ImpOs            `json:"imp"`        // Impostos
	InfCTeNorm InfCTeNormOs     `json:"infCTeNorm"` // Informações do CTe Normal
	InfCteComp InfCteComp       `json:"infCteComp"` // Informações de complemento do CTe
	InfCteAnu  InfCteAnu       `json:"infCteAnu"`  // Informações de anulação do CTe
	AutXml     []AutXML         `json:"autXml"`     // Autorização XML
	InfRespTec InfRespTec       `json:"infRespTec"` // Informações do responsável técnico
}

// VersaoServico define a versão do serviço (defina conforme necessário).
type VersaoServico string

// IdeOs define a estrutura da identificação do CTe (defina conforme necessário).
type IdeOs struct {
	// Adicione os campos relevantes aqui
}

// ComplOs define a estrutura de complementos (defina conforme necessário).
type ComplOs struct {
	// Adicione os campos relevantes aqui
}

// EmitOs define a estrutura do emitente (defina conforme necessário).
type EmitOs struct {
	// Adicione os campos relevantes aqui
}

// TomaOs define a estrutura do tomador (defina conforme necessário).
type TomaOs struct {
	// Adicione os campos relevantes aqui
}

// VPrestOs define a estrutura dos valores da prestação (defina conforme necessário).
type VPrestOs struct {
	// Adicione os campos relevantes aqui
}

// ImpOs define a estrutura dos impostos (defina conforme necessário).
type ImpOs struct {
	// Adicione os campos relevantes aqui
}

// InfCTeNormOs define a estrutura das informações do CTe Normal (defina conforme necessário).
type InfCTeNormOs struct {
	// Adicione os campos relevantes aqui
}

// InfCteComp define a estrutura das informações de complemento do CTe (defina conforme necessário).
type InfCteComp struct {
	// Adicione os campos relevantes aqui
}

// InfCteAnu define a estrutura das informações de anulação do CTe (defina conforme necessário).
type InfCteAnu struct {
	// Adicione os campos relevantes aqui
}

// AutXML define a estrutura da autorização XML (defina conforme necessário).
type AutXML struct {
	// Adicione os campos relevantes aqui
}

// InfRespTec define a estrutura das informações do responsável técnico (defina conforme necessário).
type InfRespTec struct {
	// Adicione os campos relevantes aqui
}
