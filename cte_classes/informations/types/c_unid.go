package types

// cUnid representa as unidades de medida.
type cUnid int

const (
	// M3 representa metros cúbicos.
	M3 cUnid = 0 // "00"

	// KG representa quilogramas.
 KG cUnid = 1 // "01"

	// TON representa toneladas.
	TON cUnid = 2 // "02"

	// UNIDADE representa unidades.
	UNIDADE cUnid = 3 // "03"

	// LITROS representa litros.
	LITROS cUnid = 4 // "04"

	// MMBTU representa megawatts térmicos britânicos.
	MMBTU cUnid = 5 // "05"
)
