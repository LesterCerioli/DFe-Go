package inf_cte_others_services_normal

type InfCteSubOs struct {
	ChCte     string   `xml:"chCte"`      // Chave do CT-e
	RefCteAnu string   `xml:"refCteAnu"`  // Referência ao CT-e anulado
	TomaICMS  TomaICMS `xml:"tomaICMS"`   // Tomador do ICMS
}
