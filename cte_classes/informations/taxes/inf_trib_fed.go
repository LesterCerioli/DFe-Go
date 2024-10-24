package taxes

import "math"

// InfTribFed representa as informações de tributos federais.
type InfTribFed struct {
	vPis    *float64 // Valor do PIS (pode ser nulo)
	vCofins *float64 // Valor do COFINS (pode ser nulo)
	vIr     *float64 // Valor do IR (pode ser nulo)
	vInss   *float64 // Valor do INSS (pode ser nulo)
	vCsll   *float64 // Valor da CSLL (pode ser nulo)
}

// VPIS retorna o valor do PIS com arredondamento.
func (itf *InfTribFed) VPIS() *float64 {
	if itf.vPis != nil {
		val := round(*itf.vPis, 2)
		return &val
	}
	return nil
}

// SetVPIS define o valor do PIS com arredondamento.
func (itf *InfTribFed) SetVPIS(value *float64) {
	if value != nil {
		val := round(*value, 2)
		itf.vPis = &val
	} else {
		itf.vPis = nil
	}
}

// VPISSpecified verifica se vPis possui valor.
func (itf *InfTribFed) VPISSpecified() bool {
	return itf.vPis != nil
}

// VCOFINS retorna o valor do COFINS com arredondamento.
func (itf *InfTribFed) VCOFINS() *float64 {
	if itf.vCofins != nil {
		val := round(*itf.vCofins, 2)
		return &val
	}
	return nil
}

// SetVCOFINS define o valor do COFINS com arredondamento.
func (itf *InfTribFed) SetVCOFINS(value *float64) {
	if value != nil {
		val := round(*value, 2)
		itf.vCofins = &val
	} else {
		itf.vCofins = nil
	}
}

// VCOFINSSpecified verifica se vCofins possui valor.
func (itf *InfTribFed) VCOFINSSpecified() bool {
	return itf.vCofins != nil
}

// VIR retorna o valor do IR com arredondamento.
func (itf *InfTribFed) VIR() *float64 {
	if itf.vIr != nil {
		val := round(*itf.vIr, 2)
		return &val
	}
	return nil
}

// SetVIR define o valor do IR com arredondamento.
func (itf *InfTribFed) SetVIR(value *float64) {
	if value != nil {
		val := round(*value, 2)
		itf.vIr = &val
	} else {
		itf.vIr = nil
	}
}

// VIRSpecified verifica se vIr possui valor.
func (itf *InfTribFed) VIRSpecified() bool {
	return itf.vIr != nil
}

// VINSS retorna o valor do INSS com arredondamento.
func (itf *InfTribFed) VINSS() *float64 {
	if itf.vInss != nil {
		val := round(*itf.vInss, 2)
		return &val
	}
	return nil
}

// SetVINSS define o valor do INSS com arredondamento.
func (itf *InfTribFed) SetVINSS(value *float64) {
	if value != nil {
		val := round(*value, 2)
		itf.vInss = &val
	} else {
		itf.vInss = nil
	}
}

// VINSSSpecified verifica se vInss possui valor.
func (itf *InfTribFed) VINSSSpecified() bool {
	return itf.vInss != nil
}

// VCSLL retorna o valor da CSLL com arredondamento.
func (itf *InfTribFed) VCSLL() *float64 {
	if itf.vCsll != nil {
		val := round(*itf.vCsll, 2)
		return &val
	}
	return nil
}

// SetVCSLL define o valor da CSLL com arredondamento.
func (itf *InfTribFed) SetVCSLL(value *float64) {
	if value != nil {
		val := round(*value, 2)
		itf.vCsll = &val
	} else {
		itf.vCsll = nil
	}
}

// VCSLLSpecified verifica se vCsll possui valor.
func (itf *InfTribFed) VCSLLSpecified() bool {
	return itf.vCsll != nil
}

// round arredonda um número para um número específico de casas decimais
func round(val float64, places int) float64 {
	p := math.Pow(10, float64(places))
	return math.Round(val*p) / p
}
