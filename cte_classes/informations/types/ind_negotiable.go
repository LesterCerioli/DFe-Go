package types

// indNegociavel representa o indicador de negociabilidade.
type indNegociavel int

const (
	// NaoNegociavel indica que não é negociável.
	NaoNegociavel indNegociavel = 0 // "0"

	// Negociavel indica que é negociável.
	Negociavel indNegociavel = 1 // "1"
)
