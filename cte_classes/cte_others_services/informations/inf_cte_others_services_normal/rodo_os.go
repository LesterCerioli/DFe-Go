package inf_cte_others_services_normal

// RodoOS é a estrutura equivalente em Go da classe C# rodoOS, que herda de ContainerModal
type RodoOS struct {
	TAF           string    `xml:"TAF"`
	NroRegEstadual string   `xml:"NroRegEstadual"`
	Veic          *VeicOs   `xml:"veic"`
	InfFretamento *InfFretamento `xml:"infFretamento"`
}

// VeicOs e InfFretamento são estruturas adicionais que precisam ser definidas
type VeicOs struct {
	// Definição dos campos do veículo (a ser ajustada com base na classe C# original)
}

type InfFretamento struct {
	// Definição dos campos de infFretamento (a ser ajustada com base na classe C# original)
}
