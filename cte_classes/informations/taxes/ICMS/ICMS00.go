package ICMS

import (
	"CTe/Classes/Informacoes/Impostos/ICMS/Tipos"
	"CTe/Classes/Informacoes/Tipos"
	"DFe/Classes"
)

// ICMS00 representa a classe ICMS do tipo 00.
type ICMS00 struct {
	CST   Tipos.CST       // Código de Situação Tributária
	vBC   decimal.Decimal // Valor da Base de Cálculo
	pICMS decimal.Decimal // Alíquota do ICMS
	vICMS decimal.Decimal // Valor do ICMS
}

// NewICMS00 cria uma nova instância de ICMS00 com CST definido.
func NewICMS00() *ICMS00 {
	return &ICMS00{
		CST: Tipos.ICMS00,
	}
}

// VBC retorna o valor da base de cálculo arredondado.
func (icms *ICMS00) VBC() decimal.Decimal {
	return icms.vBC.Arredondar(2)
}

// SetVBC define o valor da base de cálculo arredondado.
func (icms *ICMS00) SetVBC(value decimal.Decimal) {
	icms.vBC = value.Arredondar(2)
}

// PICMS retorna a alíquota do ICMS arredondada.
func (icms *ICMS00) PICMS() decimal.Decimal {
	return icms.pICMS.Arredondar(2)
}

// SetPICMS define a alíquota do ICMS arredondada.
func (icms *ICMS00) SetPICMS(value decimal.Decimal) {
	icms.pICMS = value.Arredondar(2)
}

// VICMS retorna o valor do ICMS arredondado.
func (icms *ICMS00) VICMS() decimal.Decimal {
	return icms.vICMS.Arredondar(2)
}

// SetVICMS define o valor do ICMS arredondado.
func (icms *ICMS00) SetVICMS(value decimal.Decimal) {
	icms.vICMS = value.Arredondar(2)
}
