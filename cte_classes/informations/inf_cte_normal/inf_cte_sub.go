package inf_cte_normal

// InfCteSub representa as informações de sub-CTe.
type InfCteSub struct {
	ChCte         string      `xml:"chCte"`         // Chave do CTe
	RefCteAnu     string      `xml:"refCteAnu"`     // Referência ao CTe anulado (Versão 3.0)
	TomaICMS      TomaICMS    `xml:"tomaICMS"`      // Informações de ICMS
	IndAlteraToma *byte       `xml:"indAlteraToma"` // Indica alteração do tomador (opcional)
	TomaNaoICMS   TomaNaoICMS `xml:"tomaNaoICMS"`   // Informações de não ICMS (Versão 2.0)
}

// IndAlteraTomaSpecified indica se IndAlteraToma está definido.
func (ics *InfCteSub) IndAlteraTomaSpecified() bool {
	return ics.IndAlteraToma != nil
}
