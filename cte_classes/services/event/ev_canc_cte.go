package event

import "encoding/xml"

// EvCancCTe representa o evento de Cancelamento (evCancCTe)
type EvCancCTe struct {
	XMLName    xml.Name `xml:"evCancCTe"`
	DescEvento string   `xml:"descEvento"`
	NProt      string   `xml:"nProt"`
	XJust      string   `xml:"xJust"`
}

// NewEvCancCTe cria uma nova instância de EvCancCTe com valores padrão
func NewEvCancCTe() *EvCancCTe {
	return &EvCancCTe{
		DescEvento: "Cancelamento",
	}
}
