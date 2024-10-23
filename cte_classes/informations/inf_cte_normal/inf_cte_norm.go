package inf_cte_normal

import "encoding/xml"

// InfCTeNorm representa as informações normais do CTe.
type InfCTeNorm struct {
	InfServico     InfServico     `xml:"infServico"`
	InfQcteOs      InfQcteOs      `xml:"infQcteOs"`
	InfDocRef      []InfDocRef    `xml:"infDocRef"`      // Lista de documentos de referência
	InfCarga       InfCarga       `xml:"infCarga"`       // Informações sobre a carga
	InfDoc         InfDoc         `xml:"infDoc"`         // Informações do documento
	DocAnt         DocAnt         `xml:"docAnt"`         // Documentos anteriores
	Seg            []Seg          `xml:"seg"`            // Seguros (Versão 2.0)
	InfModal       InfModal       `xml:"infModal"`       // Modal de transporte
	Peri           []Peri         `xml:"peri"`           // Períodos (Versão 2.0)
	VeicNovos      []VeicNovos    `xml:"veicNovos"`      // Veículos novos
	Cobr           Cobr           `xml:"cobr"`           // Cobranças
	InfCteSub      InfCteSub      `xml:"infCteSub"`      // Informações do CTe Sub
	InfGlobalizado InfGlobalizado `xml:"infGlobalizado"` // Globalização (Versão 3.0)
	InfServVinc    InfServVinc    `xml:"infServVinc"`    // Serviço vinculado (Versão 3.0)
}
