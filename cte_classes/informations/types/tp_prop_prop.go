package types

// tpPropProp representa os tipos de propriedade de veículo.
type tpPropProp int

const (
	// TACAgregado indica um tipo agregado.
	TACAgregado tpPropProp = 0

	// TACIndependente indica um tipo independente.
	TACIndependente tpPropProp = 1

	// Outros indica outros tipos.
	Outros tpPropProp = 2
)
