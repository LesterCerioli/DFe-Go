package taxes

// Imp representa as informações de impostos.
type Imp struct {
	ICMS       *TributacaoICMS // ICMS, um ponteiro para o tipo ICMS
	vTotTrib   *float64        // Valor total dos tributos (pode ser nulo)
	InfAdFisco string          // Informação adicional para a Receita Federal
	ICMSUFFim  *ICMSUFFim      // ICMS UF fim, um ponteiro para o tipo ICMSUFFim
	InfTribFed *InfTribFed     // Informações de tributos federais
}

// VTOTTrib retorna o valor total dos tributos com arredondamento
func (i *Imp) VTOTTrib() *float64 {
	if i.vTotTrib != nil {
		val := round(*i.vTotTrib, 2)
		return &val
	}
	return nil
}

// SetVTOTTrib define o valor total dos tributos com arredondamento
func (i *Imp) SetVTOTTrib(value *float64) {
	if value != nil {
		val := round(*value, 2)
		i.vTotTrib = &val
	} else {
		i.vTotTrib = nil
	}
}

// VTOTTribSpecified retorna se vTotTrib possui valor
func (i *Imp) VTOTTribSpecified() bool {
	return i.vTotTrib != nil
}

// round arredonda um número para um número específico de casas decimais
func round(val float64, places int) float64 {
	p := math.Pow(10, float64(places))
	return math.Round(val*p) / p
}
