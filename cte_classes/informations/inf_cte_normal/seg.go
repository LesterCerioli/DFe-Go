package inf_cte_normal

import "github.com/seu_usuario/seu_repositorio/dfe/classes"

// Seg representa a estrutura de seguro.
type Seg struct {
	RespSeg classes.RespSeg `json:"respSeg" xml:"respSeg"`
	XSeg    string          `json:"xSeg" xml:"xSeg"`
	NApol   string          `json:"nApol" xml:"nApol"`
	NAver   string          `json:"nAver" xml:"nAver"`
	VCarga  *float64        `json:"vCarga" xml:"vCarga,omitempty"`
}

// SetVCarga arredonda e define o valor de vCarga.
func (s *Seg) SetVCarga(value float64) {
	// Supondo que haja uma função Arredondar definida em algum lugar no seu código.
	s.VCarga = round(value, 2)
}

// GetVCarga retorna o valor arredondado de vCarga.
func (s *Seg) GetVCarga() float64 {
	if s.VCarga != nil {
		return *s.VCarga
	}
	return 0.0
}

// VCargaSpecified retorna se vCarga está definido.
func (s *Seg) VCargaSpecified() bool {
	return s.VCarga != nil
}

// Função de arredondamento
func round(value float64, places int) *float64 {
	// Adiciona a lógica de arredondamento aqui
	// Retorna o valor arredondado como ponteiro
}
