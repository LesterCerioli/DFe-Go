package types

// retira representa a opção de retirada.
type retira string

const (
	// Sim indica que a retirada é permitida.
	Sim retira = "0" // 0 - Sim
	// Nao indica que a retirada não é permitida.
	Nao retira = "1" // 1 - Não
)
