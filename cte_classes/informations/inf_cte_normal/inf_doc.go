package inf_cte_normal

// InfDoc representa informações de documentos.
type InfDoc struct {
	InfNF     []InfNF     `xml:"infNF"`     // Lista de informações de Nota Fiscal
	InfNFe    []InfNFe    `xml:"infNFe"`    // Lista de informações de Nota Fiscal Eletrônica
	InfOutros []InfOutros `xml:"infOutros"` // Lista de outras informações
	NCont     string      `xml:"nCont"`     // Número do Contêiner
	DPrev     string      `xml:"dPrev"`     // Data Prevista
}
