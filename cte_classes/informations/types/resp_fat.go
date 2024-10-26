package types

// respFat representa as responsabilidades de faturamento.
type respFat string

const (
	// FerroviaDeOrigem representa a ferrovia de origem.
	FerroviaDeOrigem respFat = "1" // 1 - Ferrovia de origem
	// FerroviaDeDestino representa a ferrovia de destino.
	FerroviaDeDestino respFat = "2" // 2 - Ferrovia de destino
)
