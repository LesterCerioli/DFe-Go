package inf_modals

// ValePed representa as informações de um vale-pedágio.
type vale_ped struct {
	CNPJForn string          `xml:"CNPJForn"`
	NCompra  string          `xml:"nCompra"`
	CNPJPg   string          `xml:"CNPJPg"`
	VValePed decimal.Decimal `xml:"vValePed"`
}
