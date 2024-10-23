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
