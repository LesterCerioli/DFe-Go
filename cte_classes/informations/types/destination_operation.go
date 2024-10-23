package types

// DestinoOperacao identifica o local de destino da operação.
type DestinoOperacao string

const (
	// DoInterna representa uma operação interna.
	DoInterna DestinoOperacao = "1"
	// DoInterestadual representa uma operação interestadual.
	DoInterestadual DestinoOperacao = "2"
	// DoExterior representa uma operação para o exterior.
	DoExterior DestinoOperacao = "3"
)
