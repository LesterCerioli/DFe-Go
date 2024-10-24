package uselessness

import (
	"encoding/xml"
)

// InfInutEnv representa as informações para a inutilização
type InfInutEnv struct {
	XServ    string         `xml:"xServ"`    // DP06 - Serviço solicitado: "INUTILIZAR"
	Id       string         `xml:"Id,attr"`  // DP04 - Identificador da TAG a ser assinada
	TpAmb    TipoAmbiente   `xml:"tpAmb"`    // DP05 - Identificação do Ambiente
	CUF      Estado         `xml:"cUF"`      // DP07 - Código da UF do solicitante
	Ano      int            `xml:"ano"`      // DP08 - Ano de inutilização da numeração
	CNPJ     string         `xml:"CNPJ"`     // DP09 - CNPJ do emitente
	Mod      ModeloDocumento `xml:"mod"`      // DP10 - Modelo da NF-e (55)
	Serie    short          `xml:"serie"`    // DP11 - Série da NF-e
	NCTIni   long           `xml:"nCTIni"`   // DP12 - Número da NF-e inicial a ser inutilizada
	NCTFin   long           `xml:"nCTFin"`   // DP13 - Número da NF-e final a ser inutilizada
	XJust    string         `xml:"xJust"`    // DP14 - Justificativa do pedido de inutilização
}

// NewInfInutEnv cria uma nova instância de InfInutEnv com o serviço inicializado
func NewInfInutEnv() *InfInutEnv {
	return &InfInutEnv{
		XServ: "INUTILIZAR",
	}
}
