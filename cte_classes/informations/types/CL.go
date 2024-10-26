package types

// CL representa o tipo de tarifa.
type CL string

const (
	CLTarifaMinima    CL = "M" // Tarifa Mínima
	CLTarifaGeral     CL = "G" // Tarifa Geral
	CLTarifaEspecifica CL = "E" // Tarifa Específica
)
