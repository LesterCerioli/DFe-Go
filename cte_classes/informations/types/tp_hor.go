package types

// tpHor representa os tipos de horário.
type tpHor int

const (
	// SemHoraDefinida representa a ausência de hora definida.
	SemHoraDefinida tpHor = iota // 0

	// NoHorario representa um horário específico.
	NoHorario // 1

	// AteOHorario representa até o horário especificado.
	AteOHorario // 2

	// ApartirDoHorario representa a partir do horário especificado.
	ApartirDoHorario // 3

	// NoIntervaloDeTempo representa um intervalo de tempo específico.
	NoIntervaloDeTempo // 4
)
