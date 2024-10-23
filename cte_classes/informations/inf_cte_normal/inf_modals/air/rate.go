package air

import "CTe.Classes.Informacoes.Tipos"

// Tarifa representa as informações sobre a tarifa de carga aérea.
type Tarifa struct {
	CL     Tipos.CL `json:"CL" xml:"CL"`
	CTar   string   `json:"cTar" xml:"cTar"`
	vTar   float64  `json:"vTar" xml:"vTar"`
}

// SetVTar arredonda o valor da tarifa para 2 casas decimais.
func (t *Tarifa) SetVTar(value float64) {
	t.vTar = round(value, 2)
}

// GetVTar retorna o valor da tarifa arredondado para 2 casas decimais.
func (t *Tarifa) GetVTar() float64 {
	return round(t.vTar, 2)
}

// round arredonda um número para um número específico de casas decimais.
func round(value float64, places int) float64 {
	if places < 0 {
		return value
	}

	p := math.Pow(10, float64(places))
	return math.Floor(value*p+0.5) / p
}
