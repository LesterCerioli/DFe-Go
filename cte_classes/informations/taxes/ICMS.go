package taxes

import (
	"CTe/Classes/Informacoes/Impostos/ICMS"
	"encoding/xml"
)

// ICMS representa a estrutura que contém informações sobre ICMS.
type ICMS struct {
	XMLName  xml.Name   `xml:"ICMS"`
	TipoICMS ICMSBasico `xml:",any"` // Utiliza o tipo 'any' para suportar diferentes tipos de ICMS
}

// Define o tipo ICMSBasico para suportar as classes ICMS00, ICMS20, ICMS45, ICMS60, ICMS90, ICMSOutraUF e ICMSSN.
type ICMSBasico interface {
	// Aqui você pode definir métodos que devem ser implementados por tipos de ICMS, se necessário.
}
