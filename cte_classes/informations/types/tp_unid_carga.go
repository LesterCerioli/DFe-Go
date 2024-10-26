package types

// tpUnidCarga representa os tipos de unidade de carga.
type tpUnidCarga int

const (
	// Container representa uma unidade de carga do tipo container.
	Container tpUnidCarga = 1

	// ULD representa uma unidade de carga do tipo ULD (Unit Load Device).
	ULD tpUnidCarga = 2

	// Pallet representa uma unidade de carga do tipo pallet.
	Pallet tpUnidCarga = 3

	// Outros representa outras unidades de carga.
	Outros tpUnidCarga = 4
)
