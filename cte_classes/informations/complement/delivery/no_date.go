package delivery

import (
	"ctedos/informacoes/tipos"
)

// SemData representa uma estrutura com período sem data definida.
type SemData struct {
	TpPer tipos.TpPer `json:"tpPer"`
}

// NewSemData cria uma nova instância de SemData com o valor padrão de tpPer como SemDataDefinida.
func NewSemData() *SemData {
	return &SemData{
		TpPer: tipos.SemDataDefinida,
	}
}
