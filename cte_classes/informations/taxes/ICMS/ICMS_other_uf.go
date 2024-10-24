package ICMS

import (
	"CTe/Classes/Informacoes/Impostos/ICMS/Tipos"
	"DFe/Classes"
)

// ICMSOutraUF representa a classe ICMS para outra unidade federativa.
type ICMSOutraUF struct {
	CST           Tipos.CST        // Código de Situação Tributária
	pRedBcOutraUf *decimal.Decimal // Percentual de redução da base de cálculo
	vBcOutraUf    decimal.Decimal  // Valor da base de cálculo
	pIcmsOutraUf  decimal.Decimal  // Percentual de ICMS
	vIcmsOutraUf  decimal.Decimal  // Valor do ICMS
}

// NewICMSOutraUF cria uma nova instância de ICMSOutraUF e inicializa CST.
func NewICMSOutraUF() *ICMSOutraUF {
	return &ICMSOutraUF{
		CST: Tipos.ICMS90,
	}
}

// SetPRedBCOutraUF define o percentual de redução da base de cálculo para outra UF.
func (i *ICMSOutraUF) SetPRedBCOutraUF(value *decimal.Decimal) {
	i.pRedBcOutraUf = value
}

// PRedBCOutraUF retorna o percentual de redução da base de cálculo para outra UF.
func (i *ICMSOutraUF) PRedBCOutraUF() *decimal.Decimal {
	return i.pRedBcOutraUf
}

// PRedBCOutraUFSpecified retorna true se pRedBcOutraUf está definido.
func (i *ICMSOutraUF) PRedBCOutraUFSpecified() bool {
	return i.pRedBcOutraUf != nil
}

// SetVBCOutraUF define o valor da base de cálculo para outra UF.
func (i *ICMSOutraUF) SetVBCOutraUF(value decimal.Decimal) {
	i.vBcOutraUf = value
}

// VBCOutraUF retorna o valor da base de cálculo para outra UF.
func (i *ICMSOutraUF) VBCOutraUF() decimal.Decimal {
	return i.vBcOutraUf
}

// SetPICMSOutraUF define o percentual de ICMS para outra UF.
func (i *ICMSOutraUF) SetPICMSOutraUF(value decimal.Decimal) {
	i.pIcmsOutraUf = value
}

// PICMSOutraUF retorna o percentual de ICMS para outra UF.
func (i *ICMSOutraUF) PICMSOutraUF() decimal.Decimal {
	return i.pIcmsOutraUf
}

// SetVICMSOutraUF define o valor do ICMS para outra UF.
func (i *ICMSOutraUF) SetVICMSOutraUF(value decimal.Decimal) {
	i.vIcmsOutraUf = value
}

// VICMSOutraUF retorna o valor do ICMS para outra UF.
func (i *ICMSOutraUF) VICMSOutraUF() decimal.Decimal {
	return i.vIcmsOutraUf
}
