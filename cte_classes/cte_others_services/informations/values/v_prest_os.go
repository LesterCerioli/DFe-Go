package values

import (
	"github.com/shopspring/decimal"
)

// Comp representa a estrutura de complemento (defina conforme necessário).
type Comp struct {
	// Adicione os campos relevantes aqui
}

// VPrestOs representa a estrutura equivalente em Go da classe C# vPrestOs.
type VPrestOs struct {
	vTPrest decimal.Decimal `json:"vTPrest" xml:"vTPrest"` // Valor Total do Prestação
	vRec    decimal.Decimal `json:"vRec" xml:"vRec"`       // Valor da Recepção
	Comp    []Comp          `json:"Comp" xml:"Comp"`       // Lista de complementos
}

// SetVTPrest define o valor total da prestação arredondado.
func (v *VPrestOs) SetVTPrest(value decimal.Decimal) {
	v.vTPrest = value.Round(2)
}

// GetVTPrest retorna o valor total da prestação arredondado.
func (v *VPrestOs) GetVTPrest() decimal.Decimal {
	return v.vTPrest.Round(2)
}

// SetVRec define o valor da recepção arredondado.
func (v *VPrestOs) SetVRec(value decimal.Decimal) {
	v.vRec = value.Round(2)
}

// GetVRec retorna o valor da recepção arredondado.
func (v *VPrestOs) GetVRec() decimal.Decimal {
	return v.vRec.Round(2)
}
