package ICMS

import (
	"CTe/Classes/Informacoes/Impostos/ICMS/Tipos"
)

// ICMSSN representa a classe ICMS Substituição Normal.
type ICMSSN struct {
	CST          *Tipos.CST  // Código de Situação Tributária, usando ponteiro para permitir valor nulo.
	CSTSpecified bool        // Indica se CST está especificado.
	indSN        Tipos.indSN // Indicador de Substituição Normal
}

// NewICMSSN cria uma nova instância de ICMSSN e inicializa o indSN.
func NewICMSSN() *ICMSSN {
	return &ICMSSN{
		indSN: Tipos.Sim,
	}
}

// SetCST define o valor do CST, permitindo valores nulos.
func (i *ICMSSN) SetCST(value *Tipos.CST) {
	i.CST = value
	i.CSTSpecified = value != nil
}
