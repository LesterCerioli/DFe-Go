package complement

// Fluxo representa o fluxo das informações complementares de um CTe.
type flow struct {
	XOrig string `xml:"xOrig"` // Origem
	Pass  []Pass `xml:"pass"`  // Lista de passagens
	XDest string `xml:"xDest"` // Destino
	XRota string `xml:"xRota"` // Rota
}
