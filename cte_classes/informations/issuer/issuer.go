package issuer

// Emit representa as informações do emitente.
type Emit struct {
	CNPJ      string    `json:"CNPJ"`          // CNPJ do emitente
	IE        string    `json:"IE"`            // Inscrição Estadual
	IEST      string    `json:"IEST"`          // Inscrição Estadual do Substituto Tributário (não é obrigatório na versão 3.00)
	XNome     string    `json:"xNome"`         // Nome do emitente
	XFant     string    `json:"xFant"`         // Nome fantasia do emitente
	EnderEmit EnderEmit `json:"enderEmit"`     // Endereço do emitente
	CRT       *CRT      `json:"CRT,omitempty"` // Classificação do Regime Tributário (obrigatório na versão 4.00)
}

// EnderEmit representa o endereço do emitente (defina conforme necessário).
type EnderEmit struct {
	// Defina os campos apropriados para o endereço do emitente aqui
}

// CRT representa a classificação do regime tributário.
type CRT int

const (
	// Defina os valores possíveis para CRT aqui
	CRT1 CRT = iota // Exemplo de um valor
	CRT2            // Outro exemplo de valor
)

// CRTSpecified retorna se CRT está definido.
func (e *Emit) CRTSpecified() bool {
	return e.CRT != nil
}
