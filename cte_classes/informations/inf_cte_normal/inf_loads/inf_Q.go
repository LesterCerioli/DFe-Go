package inf_loads

import "CTe/Classes/Informacoes/Tipos"

// InfQ representa informações sobre a carga.
type InfQ struct {
	CUnid  Tipos.CUnid `json:"cUnid"`  // Unidade de medida da carga
	TpMed  string       `json:"tpMed"`  // Tipo de medida
	qCarga float64      `json:"qCarga"` // Quantidade de carga
}

// SetQCarga define a quantidade de carga com arredondamento
func (inf *InfQ) SetQCarga(value float64) {
	inf.qCarga = arredondar(value, 4)
}

// GetQCarga retorna a quantidade de carga arredondada
func (inf *InfQ) GetQCarga() float64 {
	return arredondar(inf.qCarga, 4)
}

// arredondar é uma função auxiliar para arredondar valores para um número específico de casas decimais
func arredondar(val float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	return math.Round(val*pow) / pow
}
