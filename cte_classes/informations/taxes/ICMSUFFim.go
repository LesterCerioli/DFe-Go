package taxes

import "DFe/Classes"

// ICMSUFFim representa a estrutura que contém informações sobre ICMS UF fim.
type ICMSUFFim struct {
	vBcufFim       float64  // Valor BC UF fim
	pFcpufFim      float64  // Percentual FCP UF fim
	pIcmsufFim     float64  // Percentual ICMS UF fim
	pIcmsInter     float64  // Percentual ICMS interestadual
	pIcmsInterPart *float64 // Percentual ICMS interestadual parte
	vFcpufFim      float64  // Valor FCP UF fim
	vIcmsufFim     float64  // Valor ICMS UF fim
	vIcmsufIni     float64  // Valor ICMS UF inicial
}

// VBCUFFim retorna o valor de BC UF fim com arredondamento
func (i *ICMSUFFim) VBCUFFim() float64 {
	return round(i.vBcufFim, 2)
}

// SetVBCUFFim define o valor de BC UF fim com arredondamento
func (i *ICMSUFFim) SetVBCUFFim(value float64) {
	i.vBcufFim = round(value, 2)
}

// PFCPUFFim retorna o percentual FCP UF fim com arredondamento
func (i *ICMSUFFim) PFCPUFFim() float64 {
	return round(i.pFcpufFim, 2)
}

// SetPFCPUFFim define o percentual FCP UF fim com arredondamento
func (i *ICMSUFFim) SetPFCPUFFim(value float64) {
	i.pFcpufFim = round(value, 2)
}

// PICMSUFFim retorna o percentual ICMS UF fim com arredondamento
func (i *ICMSUFFim) PICMSUFFim() float64 {
	return round(i.pIcmsufFim, 2)
}

// SetPICMSUFFim define o percentual ICMS UF fim com arredondamento
func (i *ICMSUFFim) SetPICMSUFFim(value float64) {
	i.pIcmsufFim = round(value, 2)
}

// PICMSInter retorna o percentual ICMS interestadual com arredondamento
func (i *ICMSUFFim) PICMSInter() float64 {
	return round(i.pIcmsInter, 2)
}

// SetPICMSInter define o percentual ICMS interestadual com arredondamento
func (i *ICMSUFFim) SetPICMSInter(value float64) {
	i.pIcmsInter = round(value, 2)
}

// PICMSInterPart retorna o percentual ICMS interestadual parte com arredondamento
func (i *ICMSUFFim) PICMSInterPart() *float64 {
	if i.pIcmsInterPart != nil {
		val := round(*i.pIcmsInterPart, 2)
		return &val
	}
	return nil
}

// SetPICMSInterPart define o percentual ICMS interestadual parte com arredondamento
func (i *ICMSUFFim) SetPICMSInterPart(value *float64) {
	if value != nil {
		val := round(*value, 2)
		i.pIcmsInterPart = &val
	} else {
		i.pIcmsInterPart = nil
	}
}

// PFCPUFFim retorna o valor de FCP UF fim com arredondamento
func (i *ICMSUFFim) VFCPUFFim() float64 {
	return round(i.vFcpufFim, 2)
}

// SetVFCPUFFim define o valor de FCP UF fim com arredondamento
func (i *ICMSUFFim) SetVFCPUFFim(value float64) {
	i.vFcpufFim = round(value, 2)
}

// VICMSUFFim retorna o valor de ICMS UF fim com arredondamento
func (i *ICMSUFFim) VICMSUFFim() float64 {
	return round(i.vIcmsufFim, 2)
}

// SetVICMSUFFim define o valor de ICMS UF fim com arredondamento
func (i *ICMSUFFim) SetVICMSUFFim(value float64) {
	i.vIcmsufFim = round(value, 2)
}

// VICMSUFIni retorna o valor de ICMS UF inicial com arredondamento
func (i *ICMSUFFim) VICMSUFIni() float64 {
	return round(i.vIcmsufIni, 2)
}

// SetVICMSUFIni define o valor de ICMS UF inicial com arredondamento
func (i *ICMSUFFim) SetVICMSUFIni(value float64) {
	i.vIcmsufIni = round(value, 2)
}

// round arredonda um número para um número específico de casas decimais
func round(val float64, places int) float64 {
	p := math.Pow(10, float64(places))
	return math.Round(val*p) / p
}
