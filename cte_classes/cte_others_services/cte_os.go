package cte_others_services

import (
	"ctedos/informacoes"
	"dfe/assinatura"
	"dfe/flags"
	"dfe/utils"
)

// CTeOS representa a estrutura equivalente em Go da classe C# CTeOS.
type CTeOS struct {
	Versao     flags.VersaoServico    `json:"versao"`
	InfCte     informacoes.InfCteOS   `json:"infCte"`
	InfCTeSupl informacoes.InfCTeSupl `json:"infCTeSupl"`
	Signature  assinatura.Signature   `json:"signature"`
}

// NewCTeOS cria uma nova instância de CTeOS com a versão padrão.
func NewCTeOS() *CTeOS {
	return &CTeOS{
		Versao: flags.Versao300, // Define a versão padrão como Versao300
	}
}

// LoadXmlString carrega uma string XML e retorna uma instância de CTeOS.
func LoadXmlString(xml string) (*CTeOS, error) {
	return utils.XmlStringParaClasse[CTeOS](xml)
}

// LoadXmlArquivo carrega um arquivo XML e retorna uma instância de CTeOS.
func LoadXmlArquivo(caminhoArquivoXml string) (*CTeOS, error) {
	return utils.ArquivoXmlParaClasse[CTeOS](caminhoArquivoXml)
}
