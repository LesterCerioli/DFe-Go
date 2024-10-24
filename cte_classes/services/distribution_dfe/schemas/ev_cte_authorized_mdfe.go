package schemas

// EvCTeAutorizadoMDFe representa o evento de CTe autorizado para MDF-e.
type EvCTeAutorizadoMDFe struct {
	MDFe       MDFe   `xml:"MDFe" json:"MDFe"`
	Emit       Emit   `xml:"emit" json:"emit"`
	DescEvento string `xml:"descEvento" json:"descEvento"`
}
