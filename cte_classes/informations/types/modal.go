package types

// modal representa os tipos de modal de transporte.
type modal string

const (
	// rodoviario representa o modal rodoviário.
	rodoviario modal = "01" // "01"
	// aereo representa o modal aéreo.
	aereo modal = "02" // "02"
	// aquaviario representa o modal aquaviário.
	aquaviario modal = "03" // "03"
	// ferroviario representa o modal ferroviário.
	ferroviario modal = "04" // "04"
	// dutoviario representa o modal dutoviário.
	dutoviario modal = "05" // "05"
	// multimodal representa o modal multimodal.
	multimodal modal = "06" // "06"
)
