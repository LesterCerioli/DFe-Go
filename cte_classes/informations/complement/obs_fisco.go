package complement

// ObsFisco representa uma observação fiscal com um campo e texto associado.
type ObsFisco struct {
	XCampo string `xml:"xCampo,attr"` // Atributo XML xCampo
	XTexto string `xml:"xTexto"`      // Elemento XML xTexto
}
