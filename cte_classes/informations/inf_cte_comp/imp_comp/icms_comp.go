package imp_comp

import (
	"CTe/Classes/Informacoes/Impostos/ICMS"
)

// ICMSComp representa a composição de ICMS.
type ICMSComp struct {
	ICMS00        *ICMS.ICMS00        `xml:"ICMS00,omitempty"`        // ICMS00 pode ser nulo
	ICMS20        *ICMS.ICMS20        `xml:"ICMS20,omitempty"`        // ICMS20 pode ser nulo
	ICMS45        *ICMS.ICMS45        `xml:"ICMS45,omitempty"`        // ICMS45 pode ser nulo
	ICMS60        *ICMS.ICMS60        `xml:"ICMS60,omitempty"`        // ICMS60 pode ser nulo
	ICMS90        *ICMS.ICMS90        `xml:"ICMS90,omitempty"`        // ICMS90 pode ser nulo
	ICMSOutraUF   *ICMS.ICMSOutraUF   `xml:"ICMSOutraUF,omitempty"`   // ICMSOutraUF pode ser nulo
	ICMSSN        *ICMS.ICMSSN        `xml:"ICMSSN,omitempty"`        // ICMSSN pode ser nulo
}
