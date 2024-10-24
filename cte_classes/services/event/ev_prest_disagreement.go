package event

import "encoding/xml"

// EvPrestDesacordo representa o evento de Prestação do Serviço em Desacordo (evPrestDesacordo)
type EvPrestDesacordo struct {
	XMLName           xml.Name `xml:"evPrestDesacordo"`
	DescEvento        string   `xml:"descEvento"`
	IndDesacordoOper  string   `xml:"indDesacordoOper"`
	XObs              string   `xml:"xObs"`
}

// NewEvPrestDesacordo cria uma nova instância de EvPrestDesacordo com valores padrão
func NewEvPrestDesacordo() *EvPrestDesacordo {
	return &EvPrestDesacordo{
		DescEvento: "Prestação do Serviço em Desacordo",
	}
}
