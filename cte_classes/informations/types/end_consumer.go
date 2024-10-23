package types

// ConsumidorFinal indica operação com Consumidor final.
type ConsumidorFinal string

const (
	// CfNao representa uma operação normal (não consumidor final).
	CfNao ConsumidorFinal = "0"
	// CfConsumidorFinal representa uma operação com consumidor final.
	CfConsumidorFinal ConsumidorFinal = "1"
)
