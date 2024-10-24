package inf_cte_others_services_normal

type InfCTeNormOs struct {
	InfServico InfServico     `xml:"infServico"`  // Informações sobre o serviço
	InfDocRef  []InfDocRef    `xml:"infDocRef"`   // Lista de referências de documentos
	Seg        []SegOs        `xml:"seg"`         // Lista de seguros
	InfModal   InfModalOs     `xml:"infModal"`    // Informações do modal
	InfCteSub  InfCteSubOs    `xml:"infCteSub"`   // Informações do CT-e substituto
}
