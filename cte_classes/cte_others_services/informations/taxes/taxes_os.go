package taxes

// ImpOs representa a estrutura equivalente em Go da classe C# impOs.
type ImpOs struct {
	ICMS        ICMS         `json:"ICMS"`        // ICMS
	VTotTrib    *decimal.Decimal `json:"vTotTrib"`   // Valor Total de Tributos
	InfAdFisco  string        `json:"infAdFisco"`  // Informações Adicionais ao Fisco
	ICMSUFFim   ICMSUFFim    `json:"ICMSUFFim"`   // ICMS UF Fim
	InfTribFed  InfTribFed    `json:"infTribFed"`  // Informações de Tributação Federal
}

// SetVTotTrib define o valor total de tributos arredondado.
func (i *ImpOs) SetVTotTrib(value decimal.Decimal) {
	i.VTotTrib = value.Round(2) // Arredondar para 2 casas decimais
}

// GetVTotTrib retorna o valor total de tributos arredondado.
func (i *ImpOs) GetVTotTrib() decimal.Decimal {
	return i.VTotTrib
}

// VTOTTribSpecified indica se o valor total de tributos está especificado.
func (i *ImpOs) VTOTTribSpecified() bool {
	return i.VTotTrib != nil
}
