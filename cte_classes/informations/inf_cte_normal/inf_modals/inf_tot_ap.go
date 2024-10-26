package inf_modals

import (
	"math"
)

// InfTotAP representa a estrutura de total de produtos.
type InfTotAP struct {
	QTotProd float64 `xml:"qTotProd"`
	UniAP    string  `xml:"uniAP"`
}

// Arredondar arredonda o valor para o número de casas decimais especificadas.
func Arredondar(valor float64, casasDecimais int) float64 {
	fator := math.Pow(10, float64(casasDecimais))
	return math.Round(valor*fator) / fator
}

// SetQTotProd define o valor arredondado de QTotProd.
func (i *InfTotAP) SetQTotProd(valor float64) {
	i.QTotProd = Arredondar(valor, 4)
}
