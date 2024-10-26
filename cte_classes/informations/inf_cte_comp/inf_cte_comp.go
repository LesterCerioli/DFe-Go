package inf_cte_comp

// infCteComp representa as informações de um componente CT-e.
type infCteComp struct {
	Chave  string `xml:"chave"`  // Versão 2.0
	ChCTe  string `xml:"chCTe"`  // Versão 3.0
}

// NewInfCteComp cria uma nova instância de infCteComp.
func NewInfCteComp(chave, chCTe string) *infCteComp {
	return &infCteComp{
		Chave: chave,
		ChCTe: chCTe,
	}
}
