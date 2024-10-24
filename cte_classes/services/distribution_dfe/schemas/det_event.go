package schemas

// DetEvento representa os detalhes do evento de distribuição de documentos fiscais eletrônicos.
type DetEvento struct {
	EvCTeAutorizadoMDFe EvCTeAutorizadoMDFe `xml:"evCTeAutorizadoMDFe" json:"evCTeAutorizadoMDFe"`
	Versao              float64             `xml:"versao,attr" json:"versao"`
}
