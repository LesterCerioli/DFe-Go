package ICMS

import (
	"CTe/Classes/Informacoes/Impostos/ICMS/Tipos"
	"CTe/Classes/Informacoes/Tipos"
	"DFe/Classes"
)

// ICMS20 representa a classe ICMS do tipo 20.
type ICMS20 struct {
	CST    Tipos.CST       // Código de Situação Tributária
	pRedBC decimal.Decimal // Percentual de Redução da Base de Cálculo
	vBC    decimal.Decimal // Valor da Base de Cálculo
	pICMS  decimal.Decimal // Alíquota do ICMS
	vICMS  decimal.Decimal // Valor do ICMS
}

// NewICMS20 cria uma nova instância de ICMS20 com CST definido.
func NewICMS20() *ICMS20 {
	return &ICMS20{
		CST: Tipos.ICMS20,
	}
}

// PRedBC retorna o percentual de redução da base de cálculo arredondado.
func (icms *ICMS20) PRedBC() decimal.Decimal {
	return icms.pRedBC.Arredondar(2)
}

// SetPRedBC define o percentual de redução da base de cálculo arredondado.
func (icms *ICMS20) SetPRedBC(value decimal.Decimal) {
	icms.pRedBC = value.Arredondar(2)
}

// VBC retorna o valor da base de cálculo arredondado.
func (icms *ICMS20) VBC() decimal.Decimal {
	return icms.vBC.Arredondar(2)
}

// SetVBC define o valor da base de cálculo arredondado.
func (icms *ICMS20) SetVBC(value decimal.Decimal) {
	icms.vBC = value.Arredondar(2)
}

// PICMS retorna a alíquota do ICMS arredondada.
func (icms *ICMS20) PICMS() decimal.Decimal {
	return icms.pICMS.Arredondar(2)
}

// SetPICMS define a alíquota do ICMS arredondada.
func (icms *ICMS20) SetPICMS(value decimal.Decimal) {
	icms.pICMS = value.Arredondar(2)
}

// VICMS retorna o valor do ICMS arredondado.
func (icms *ICMS20) VICMS() decimal.Decimal {
	return icms.vICMS.Arredondar(2)
}

// SetVICMS define o valor do ICMS arredondado.
func (icms *ICMS20) SetVICMS(value decimal.Decimal) {
	icms.vICMS = value.Arredondar(2)
}
