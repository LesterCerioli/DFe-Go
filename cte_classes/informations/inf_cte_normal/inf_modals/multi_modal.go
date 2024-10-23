package inf_modals

import "CTe.Classes.Informacoes.Tipos"

// Multimodal representa a estrutura de transporte multimodal.
type Multimodal struct {
	COTM          string              `xml:"COTM"`
	IndNegociavel Tipos.IndNegociavel `xml:"indNegociavel"`
	Seg           SegMultiModal       `xml:"seg"`
}

// SegMultiModal representa as informações do seguro multimodal.
type SegMultiModal struct {
	InfSeg InfSeg `xml:"infSeg"`
	NApol  string `xml:"nApol"`
	NAver  string `xml:"nAver"`
}

// InfSeg representa as informações do seguro.
type InfSeg struct {
	XSeg string `xml:"xSeg"`
	CNPJ string `xml:"CNPJ"`
}
