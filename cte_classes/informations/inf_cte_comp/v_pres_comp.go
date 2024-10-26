package inf_cte_comp

// compComp representa um componente que contém informações de um produto ou serviço.
type compComp struct {
	XNome string  `xml:"xNome"` // Nome do componente
	VComp float64 `xml:"vComp"` // Valor do componente
}

// vPresComp representa a estrutura de valores de apresentação do componente CT-e.
type vPresComp struct {
	CompComp compComp `xml:"compComp"` // Componente do CT-e

	vTPrest float64 `xml:"vTPrest"` // Valor total do prestador
}

// NewVPresComp cria uma nova instância de vPresComp.
func NewVPresComp(compComp compComp, vTPrest float64) *vPresComp {
	return &vPresComp{
		CompComp: compComp,
		vTPrest:  vTPrest,
	}
}
