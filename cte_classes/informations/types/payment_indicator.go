package types

// IndicadorPagamento indica a forma de pagamento.
type IndicadorPagamento int

const (
	// IpVista indica pagamento à vista.
	IpVista IndicadorPagamento = 0 // pagamento à vista

	// IpPrazo indica pagamento à prazo.
	IpPrazo IndicadorPagamento = 1 // pagamento à prazo

	// IpOutras indica outros tipos de pagamento.
	IpOutras IndicadorPagamento = 2 // outros
)
