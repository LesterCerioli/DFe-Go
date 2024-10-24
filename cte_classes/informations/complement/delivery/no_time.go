package delivery

import (
	"ctedos/informacoes/tipos"
)

// SemHora representa uma estrutura com horário sem hora definida.
type SemHora struct {
	TpHor tipos.TpHor `json:"tpHor"`
}

// NewSemHora cria uma nova instância de SemHora com o valor padrão de TpHor como SemHoraDefinida.
func NewSemHora() *SemHora {
	return &SemHora{
		TpHor: tipos.SemHoraDefinida,
	}
}
