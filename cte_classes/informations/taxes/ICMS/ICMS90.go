package ICMS

import (
	"CTe/Classes/Informacoes/Impostos/ICMS/Tipos"
	"math"
)

// ICMS90 representa a classe ICMS do tipo 90.
type ICMS90 struct {
	CST    Tipos.CST // Código de Situação Tributária
	pRedBc *float64  // Percentual de Redução da Base de Cálculo
	vBc    float64   // Valor da Base de Cálculo
	pIcms  float64   // Percentual do ICMS
	vIcms  float64   // Valor do ICMS
	vCred  float64   // Valor do Crédito
}

// NewICMS90 cria uma nova instância de ICMS90 e inicializa o CST.
func NewICMS90() *ICMS90 {
	return &ICMS90{
		CST: Tipos.ICMS90,
	}
}

// PRedBC retorna o percentual de redução da base de cálculo arredondado para 2 casas decimais.
func (i *ICMS90) PRedBC() *float64 {
	if i.pRedBc != nil {
		val := *i.pRedBc
		return round(val, 2)
	}
	return nil
}

// SetPRedBC define o percentual de redução da base de cálculo arredondando para 2 casas decimais.
func (i *ICMS90) SetPRedBC(value *float64) {
	if value != nil {
		val := round(*value, 2)
		i.pRedBc = &val
	} else {
		i.pRedBc = nil
	}
}

// PRedBCSpecified verifica se o percentual de redução da base de cálculo está definido.
func (i *ICMS90) PRedBCSpecified() bool {
	return i.pRedBc != nil
}

// VBC retorna o valor da base de cálculo arredondado para 2 casas decimais.
func (i *ICMS90) VBC() float64 {
	return round(i.vBc, 2)
}

// SetVBC define o valor da base de cálculo arredondando para 2 casas decimais.
func (i *ICMS90) SetVBC(value float64) {
	i.vBc = round(value, 2)
}

// PICMS retorna o percentual do ICMS arredondado para 2 casas decimais.
func (i *ICMS90) PICMS() float64 {
	return round(i.pIcms, 2)
}

// SetPICMS define o percentual do ICMS arredondando para 2 casas decimais.
func (i *ICMS90) SetPICMS(value float64) {
	i.pIcms = round(value, 2)
}

// VICMS retorna o valor do ICMS arredondado para 2 casas decimais.
func (i *ICMS90) VICMS() float64 {
	return round(i.vIcms, 2)
}

// SetVICMS define o valor do ICMS arredondando para 2 casas decimais.
func (i *ICMS90) SetVICMS(value float64) {
	i.vIcms = round(value, 2)
}

// VCred retorna o valor do Crédito arredondado para 2 casas decimais.
func (i *ICMS90) VCred() float64 {
	return round(i.vCred, 2)
}

// SetVCred define o valor do Crédito arredondando para 2 casas decimais.
func (i *ICMS90) SetVCred(value float64) {
	i.vCred = round(value, 2)
}

// round arredonda um número para o número de casas decimais especificado.
func round(value float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(value*multiplier) / multiplier
}
