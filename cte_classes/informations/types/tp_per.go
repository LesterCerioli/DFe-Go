package types

// tpPer representa os tipos de períodos.
type tpPer int

const (
	// SemDataDefinida indica que não há data definida.
	SemDataDefinida tpPer = 0

	// NaData indica que a operação ocorre na data especificada.
	NaData tpPer = 1

	// AteAData indica que a operação ocorre até a data especificada.
	AteAData tpPer = 2

	// ApartirDaData indica que a operação ocorre a partir da data especificada.
	ApartirDaData tpPer = 3

	// NoPeriodo indica que a operação ocorre em um período específico.
	NoPeriodo tpPer = 4
)
