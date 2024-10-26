package inf_cte_normal

import "CTe/Classes" // ajuste o caminho conforme sua estrutura de pacotes

// InfQcteOs representa as informações de quantidade de carga.
type InfQcteOs struct {
	qCarga float64 // Usando float64 para decimal
}

// SetQCarga arredonda e define o valor de qCarga.
func (i *InfQcteOs) SetQCarga(value float64) {
	i.qCarga = Classes.Arredondar(value, 4) // Função de arredondamento hipotética
}

// GetQCarga retorna o valor de qCarga arredondado.
func (i *InfQcteOs) GetQCarga() float64 {
	return Classes.Arredondar(i.qCarga, 4)
}
