package inf_ctec_comp

// impComp representa as informações de imposto de um componente.
type impComp struct {
	ICMSComp ICMSComp `xml:"ICMSComp,omitempty"` // Composição de ICMS
}

// NewImpComp cria uma nova instância de impComp.
func NewImpComp(icmsComp ICMSComp) *impComp {
	return &impComp{
		ICMSComp: icmsComp,
	}
}
