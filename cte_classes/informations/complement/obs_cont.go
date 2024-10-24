package complement

// ObsCont representa uma observação com um campo e texto associado.
type ObsCont struct {
	XCampo string `xml:"xCampo,attr"` // Atributo XML xCampo
	XTexto string `xml:"xTexto"`      // Elemento XML xTexto
}
