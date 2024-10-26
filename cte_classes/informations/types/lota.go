package types

// lota representa um indicador de lotação.
type lota int

const (
	// Nao indica que não é lotado.
	Nao lota = 0 // "0"
	// Sim indica que é lotado.
	Sim lota = 1 // "1"
)
