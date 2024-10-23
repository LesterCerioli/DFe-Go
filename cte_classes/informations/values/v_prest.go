package values

import "math"

// Comp representa uma estrutura que contém informações de valores.
type Comp struct {
	xNome string  // Nome
	vComp float64 // Valor com precisão
}

// SetXNome define o nome.
func (c *Comp) SetXNome(nome string) {
	c.xNome = nome
}

// GetXNome obtém o nome.
func (c *Comp) GetXNome() string {
	return c.xNome
}

// SetVComp define o valor, arredondando para duas casas decimais.
func (c *Comp) SetVComp(valor float64) {
	c.vComp = arredondar(valor, 2)
}

// GetVComp obtém o valor arredondado.
func (c *Comp) GetVComp() float64 {
	return arredondar(c.vComp, 2)
}

// arredondar arredonda um número para o número de casas decimais especificado.
func arredondar(valor float64, casasDecimais int) float64 {
	fator := math.Pow(10, float64(casasDecimais))
	return math.Round(valor*fator) / fator
}

// VPrest representa a estrutura de valores de prestação.
type VPrest struct {
	vTPrest float64 // Valor total da prestação
	vRec    float64 // Valor recebido
	Comp    []Comp  // Lista de componentes
}

// SetVTPrest define o valor total da prestação, arredondando para duas casas decimais.
func (v *VPrest) SetVTPrest(valor float64) {
	v.vTPrest = arredondar(valor, 2)
}

// GetVTPrest obtém o valor total da prestação arredondado.
func (v *VPrest) GetVTPrest() float64 {
	return arredondar(v.vTPrest, 2)
}

// SetVRec define o valor recebido, arredondando para duas casas decimais.
func (v *VPrest) SetVRec(valor float64) {
	v.vRec = arredondar(valor, 2)
}

// GetVRec obtém o valor recebido arredondado.
func (v *VPrest) GetVRec() float64 {
	return arredondar(v.vRec, 2)
}
