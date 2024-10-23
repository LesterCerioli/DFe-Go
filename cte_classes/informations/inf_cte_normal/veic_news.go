package inf_cte_normal

import "github.com/seu_usuario/seu_repositorio/dfe/classes"

// VeicNovos representa a estrutura de um veículo novo.
type VeicNovos struct {
	Chassi string   `json:"chassi" xml:"chassi"`
	CCor   string   `json:"cCor" xml:"cCor"`
	XCor   string   `json:"xCor" xml:"xCor"`
	CMod   string   `json:"cMod" xml:"cMod"`
	VUnit  *float64 `json:"vUnit" xml:"vUnit,omitempty"`
	VFrete *float64 `json:"vFrete" xml:"vFrete,omitempty"`
}

// SetVUnit arredonda e define o valor de vUnit.
func (v *VeicNovos) SetVUnit(value float64) {
	// Supondo que haja uma função Arredondar definida em algum lugar no seu código.
	v.VUnit = round(value, 2)
}

// GetVUnit retorna o valor arredondado de vUnit.
func (v *VeicNovos) GetVUnit() float64 {
	if v.VUnit != nil {
		return *v.VUnit
	}
	return 0.0
}

// SetVFrete arredonda e define o valor de vFrete.
func (v *VeicNovos) SetVFrete(value float64) {
	// Supondo que haja uma função Arredondar definida em algum lugar no seu código.
	v.VFrete = round(value, 2)
}

// GetVFrete retorna o valor arredondado de vFrete.
func (v *VeicNovos) GetVFrete() float64 {
	if v.VFrete != nil {
		return *v.VFrete
	}
	return 0.0
}

// VCargaSpecified retorna se vUnit está definido.
func (v *VeicNovos) VUnitSpecified() bool {
	return v.VUnit != nil
}

// VFreteSpecified retorna se vFrete está definido.
func (v *VeicNovos) VFreteSpecified() bool {
	return v.VFrete != nil
}

// Função de arredondamento
func round(value float64, places int) *float64 {
	// Adicione a lógica de arredondamento aqui
	// Retorne o valor arredondado como ponteiro
}
