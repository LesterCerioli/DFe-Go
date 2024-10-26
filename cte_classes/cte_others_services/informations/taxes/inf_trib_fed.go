package taxes

import (
	"github.com/shopspring/decimal"
)

// InfTribFed representa a estrutura equivalente em Go da classe C# infTribFed.
type InfTribFed struct {
	vPis   *decimal.Decimal `json:"vPIS"`    // Valor do PIS
	vCofins *decimal.Decimal `json:"vCOFINS"` // Valor do COFINS
	vIr    *decimal.Decimal `json:"vIR"`     // Valor do IR
	vInss  *decimal.Decimal `json:"vINSS"`   // Valor do INSS
	vCsll  *decimal.Decimal `json:"vCSLL"`   // Valor da CSLL
}

// SetVPis define o valor do PIS arredondado.
func (i *InfTribFed) SetVPis(value decimal.Decimal) {
	i.vPis = &value
}

// GetVPis retorna o valor do PIS arredondado.
func (i *InfTribFed) GetVPis() *decimal.Decimal {
	if i.vPis != nil {
		return decimalPtr(i.vPis.Round(2))
	}
	return nil
}

// SetVCofins define o valor do COFINS arredondado.
func (i *InfTribFed) SetVCofins(value decimal.Decimal) {
	i.vCofins = &value
}

// GetVCofins retorna o valor do COFINS arredondado.
func (i *InfTribFed) GetVCofins() *decimal.Decimal {
	if i.vCofins != nil {
		return decimalPtr(i.vCofins.Round(2))
	}
	return nil
}

// SetVIr define o valor do IR arredondado.
func (i *InfTribFed) SetVIr(value decimal.Decimal) {
	i.vIr = &value
}

// GetVIr retorna o valor do IR arredondado.
func (i *InfTribFed) GetVIr() *decimal.Decimal {
	if i.vIr != nil {
		return decimalPtr(i.vIr.Round(2))
	}
	return nil
}

// SetVInss define o valor do INSS arredondado.
func (i *InfTribFed) SetVInss(value decimal.Decimal) {
	i.vInss = &value
}

// GetVInss retorna o valor do INSS arredondado.
func (i *InfTribFed) GetVInss() *decimal.Decimal {
	if i.vInss != nil {
		return decimalPtr(i.vInss.Round(2))
	}
	return nil
}

// SetVCsll define o valor da CSLL arredondado.
func (i *InfTribFed) SetVCsll(value decimal.Decimal) {
	i.vCsll = &value
}

// GetVCsll retorna o valor da CSLL arredondado.
func (i *InfTribFed) GetVCsll() *decimal.Decimal {
	if i.vCsll != nil {
		return decimalPtr(i.vCsll.Round(2))
	}
	return nil
}

// VPISSpecified indica se o valor do PIS está especificado.
func (i *InfTribFed) VPISSpecified() bool {
	return i.vPis != nil
}

// VCOFINSSpecified indica se o valor do COFINS está especificado.
func (i *InfTribFed) VCOFINSSpecified() bool {
	return i.vCofins != nil
}

// VIRSpecified indica se o valor do IR está especificado.
func (i *InfTribFed) VIRSpecified() bool {
	return i.vIr != nil
}

// VINSSSpecified indica se o valor do INSS está especificado.
func (i *InfTribFed) VINSSSpecified() bool {
	return i.vInss != nil
}

// VCSLLSpecified indica se o valor da CSLL está especificado.
func (i *InfTribFed) VCSLLSpecified() bool {
	return i.vCsll != nil
}

// Helper function to return a pointer to a decimal.Decimal
func decimalPtr(d decimal.Decimal) *decimal.Decimal {
	return &d
}
