package types

// CST representa as informações relativas ao ICMS.
type CST string

const (
	// ICMS00 - Tributação normal ICMS
	ICMS00 CST = "00"
	// ICMS20 - Tributação com BC reduzida do ICMS
	ICMS20 CST = "20"
	// ICMS40 - ICMS isenção
	ICMS40 CST = "40"
	// ICMS41 - ICMS não tributada
	ICMS41 CST = "41"
	// ICMS51 - ICMS diferido
	ICMS51 CST = "51"
	// ICMS60 - ICMS cobrado por substituição tributária
	ICMS60 CST = "60"
	// ICMS90 - Outros
	ICMS90 CST = "90"
)
