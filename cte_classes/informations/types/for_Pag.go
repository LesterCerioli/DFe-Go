package types

// forPag representa os tipos de forma de pagamento.
type forPag int

const (
	// Pago representa um pagamento realizado.
	Pago forPag = iota // "0"

	// Apagar representa um pagamento a ser feito.
	Apagar // "1"

	// Outros representa outras formas de pagamento.
	Outros // "2"
)
