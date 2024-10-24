package types

// ferrEmi representa os tipos de ferrovia de emissão.
type ferrEmi int

const (
	// FerroviaDeOrigem representa a ferrovia de origem.
	FerroviaDeOrigem ferrEmi = 1 // "1"

	// FerroviaDeDestino representa a ferrovia de destino.
	FerroviaDeDestino ferrEmi = 2 // "2"
)
