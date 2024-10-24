package informations

import "encoding/xml"

// InfCte representa as informações de um CTe (Conhecimento de Transporte Eletrônico).
type InfCte struct {
	XMLName    xml.Name   `xml:"infCte"`
	Versao     Versao     `xml:"versao,attr"`
	ID         string     `xml:"Id,attr"`
	Ide        Ide        `xml:"ide"`
	Compl      Compl      `xml:"compl"`
	Emit       Emit       `xml:"emit"`
	Toma       TomaCteOs  `xml:"tomaCteOs"`
	Rem        Rem        `xml:"rem"`
	Exped      Exped      `xml:"exped"`
	Receb      Receb      `xml:"receb"`
	Dest       Dest       `xml:"dest"`
	VPrest     VPrest     `xml:"vPrest"`
	Imp        Imp        `xml:"imp"`
	InfCTeNorm InfCTeNorm `xml:"infCTeNorm"`
	InfCteComp InfCteComp `xml:"infCteComp"`
	InfCteAnu  InfCteAnu  `xml:"infCteAnu"`
	AutXML     []AutXML   `xml:"autXML"`
	InfRespTec InfRespTec `xml:"infRespTec"`
}

// Defina as outras estruturas que foram referenciadas acima aqui, por exemplo:
// type Versao struct {}
// type Ide struct {}
// type Compl struct {}
// type Emit struct {}
// type TomaCteOs struct {}
// type Rem struct {}
// type Exped struct {}
// type Receb struct {}
// type Dest struct {}
// type VPrest struct {}
// type Imp struct {}
// type InfCTeNorm struct {}
// type InfCteComp struct {}
// type InfCteAnu struct {}
// type AutXML struct {}
// type InfRespTec struct {}
