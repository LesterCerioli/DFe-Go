package complement

// Compl representa as informações complementares de um CTe.
type Compl struct {
	XCaracAd  string  `xml:"xCaracAd"`  // Característica adicional
	XCaracSer string  `xml:"xCaracSer"` // Característica de série
	XEmi      string  `xml:"xEmi"`      // Emitente
	Fluxo     flow    `xml:"flow"`      // Fluxo
	Entrega   Entrega `xml:"Entrega"`   // Entrega
	OrigCalc  string  `xml:"origCalc"`  // Origem do cálculo
	DestCalc  string  `xml:"destCalc"`  // Destino do cálculo
	XObs      string  `xml:"xObs"`      // Observação

	ObsCont  []ObsCont  `xml:"ObsCont"`  // Observações de controle
	ObsFisco []ObsFisco `xml:"ObsFisco"` // Observações fiscais
}
