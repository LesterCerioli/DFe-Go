package types

// tpCTe representa os tipos de Documento Fiscal.
type tpCTe string

const (
	// Normal indica um CT-e Normal.
	Normal tpCTe = "0" // 0 - CT-e Normal
	// Complemento indica um CT-e de Complemento de Valores.
	Complemento tpCTe = "1" // 1 - CT-e de Complemento de Valores
	// Anulacao indica um CT-e de Anulação.
	Anulacao tpCTe = "2" // 2 - CT-e de Anulação
	// Substituto indica um CT-e Substituto.
	Substituto tpCTe = "3" // 3 - CT-e Substituto
)
