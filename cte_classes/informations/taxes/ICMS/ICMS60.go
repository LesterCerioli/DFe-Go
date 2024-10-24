package ICMS

import (
	"CTe/Classes/Informacoes/Impostos/ICMS/Tipos"
)

// ICMS60 representa a classe ICMS do tipo 60.
type ICMS60 struct {
	CST        Tipos.CST // Código de Situação Tributária
	vBcstRet   float64   // Valor da BC do ICMS ST Retido
	vIcmsstRet float64   // Valor do ICMS ST Retido
	pIcmsstRet float64   // Percentual do ICMS ST Retido
	vCred      float64   // Valor do Crédito
}

// NewICMS60 cria uma nova instância de ICMS60 e inicializa o CST.
func NewICMS60() *ICMS60 {
	return &ICMS60{
		CST: Tipos.ICMS60,
	}
}

// VBCSTRet retorna o valor da BC do ICMS ST Retido arredondado para 2 casas decimais.
func (i *ICMS60) VBCSTRet() float64 {
	return round(i.vBcstRet, 2)
}

// SetVBCSTRet define o valor da BC do ICMS ST Retido arredondando para 2 casas decimais.
func (i *ICMS60) SetVBCSTRet(value float64) {
	i.vBcstRet = round(value, 2)
}

// VICMSSTRet retorna o valor do ICMS ST Retido arredondado para 2 casas decimais.
func (i *ICMS60) VICMSSTRet() float64 {
	return round(i.vIcmsstRet, 2)
}

// SetVICMSSTRet define o valor do ICMS ST Retido arredondando para 2 casas decimais.
func (i *ICMS60) SetVICMSSTRet(value float64) {
	i.vIcmsstRet = round(value, 2)
}

// PICMSSTRet retorna o percentual do ICMS ST Retido arredondado para 2 casas decimais.
func (i *ICMS60) PICMSSTRet() float64 {
	return round(i.pIcmsstRet, 2)
}

// SetPICMSSTRet define o percentual do ICMS ST Retido arredondando para 2 casas decimais.
func (i *ICMS60) SetPICMSSTRet(value float64) {
	i.pIcmsstRet = round(value, 2)
}

// VCred retorna o valor do Crédito arredondado para 2 casas decimais.
func (i *ICMS60) VCred() float64 {
	return round(i.vCred, 2)
}

// SetVCred define o valor do Crédito arredondando para 2 casas decimais.
func (i *ICMS60) SetVCred(value float64) {
	i.vCred = round(value, 2)
}

// round arredonda um número para o número de casas decimais especificado.
func round(value float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(value*multiplier) / multiplier
}
