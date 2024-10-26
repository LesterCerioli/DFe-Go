package inf_cte_subs

// TomaICMS representa informações relacionadas ao ICMS.
type TomaICMS struct {
	RefNFe string  `json:"refNFe"` // Referência da nota fiscal eletrônica
	RefNF  *RefNF  `json:"refNF"`  // Referência de nota fiscal
	RefCTE string  `json:"refCte"` // Referência do CTe
}
