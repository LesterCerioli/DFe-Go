package identification

import (
	"fmt"
	"time"
	"xmlcte/constantes"
	"xmlcte/utils"
)

type IdeOs struct {
	CUF         constantes.Estado `xml:"cUF"`         // Código da UF do emitente do Documento Fiscal
	CCT         int               `xml:"-"`           // Ignorar na serialização direta
	ProxyCCT    string            `xml:"cCT"`         // Representação do cCT no formato de string
	CFOP        int               `xml:"CFOP"`        // Código Fiscal de Operações e Prestações
	NatOp       string            `xml:"natOp"`       // Natureza da Operação
	Mod         string            `xml:"mod"`         // Modelo do Documento Fiscal
	Serie       int16             `xml:"serie"`       // Série do Documento Fiscal
	NCT         int64             `xml:"nCT"`         // Número do Conhecimento de Transporte
	DhEmi       time.Time         `xml:"-"`           // Data de emissão do documento (ignorada no XML direto)
	ProxyDhEmi  string            `xml:"dhEmi"`       // Proxy para serialização da data de emissão no formato string
	TpImp       int               `xml:"tpImp"`       // Tipo de Impressão
	TpEmis      int               `xml:"tpEmis"`      // Tipo de Emissão
	CDV         int               `xml:"cDV"`         // Dígito verificador do CT-e
	TpAmb       int               `xml:"tpAmb"`       // Identificação do Ambiente
	TpCTe       int               `xml:"tpCTe"`       // Tipo do Documento Fiscal
	ProcEmi     int               `xml:"procEmi"`     // Processo de emissão
	VerProc     string            `xml:"verProc"`     // Versão do processo de emissão
	CMunEnv     int64             `xml:"cMunEnv"`     // Código do município de envio
	XMunEnv     string            `xml:"xMunEnv"`     // Nome do município de envio
	UFEnv       constantes.Estado `xml:"-"`           // UF do envio (não incluída diretamente no XML)
	ProxyUFEnv  string            `xml:"UFEnv"`       // Representação da UF no formato string
	Modal       int               `xml:"modal"`       // Modalidade do serviço
	TpServ      int               `xml:"tpServ"`      // Tipo de Serviço
	IndIEToma   int               `xml:"indIEToma"`   // Indicador de IE do tomador
	CMunIni     int64             `xml:"cMunIni"`     // Código do município de início
	XMunIni     string            `xml:"xMunIni"`     // Nome do município de início
	UFIni       constantes.Estado `xml:"-"`           // UF do início (não incluída diretamente no XML)
	ProxyUFIni  string            `xml:"UFIni"`       // Representação da UF no formato string
	CMunFim     int64             `xml:"cMunFim"`     // Código do município de fim
	XMunFim     string            `xml:"xMunFim"`     // Nome do município de fim
	UFFim       constantes.Estado `xml:"-"`           // UF do fim (não incluída diretamente no XML)
	ProxyUFFim  string            `xml:"UFFim"`       // Representação da UF no formato string
	InfPercurso []InfPercurso      `xml:"infPercurso"` // Informações do percurso
	DhCont      *time.Time        `xml:"-"`           // Data e hora do contingenciamento (opcional)
	ProxyDhCont string            `xml:"dhCont"`      // Representação da data de contingenciamento
	XJust       string            `xml:"xJust"`       // Justificativa do contingenciamento
}

// Função para formatar cCT como string
func (i *IdeOs) FormatCCT() {
	i.ProxyCCT = fmt.Sprintf("%08d", i.CCT)
}

// Função para converter DhEmi para string
func (i *IdeOs) FormatDhEmi() {
	i.ProxyDhEmi = utils.FormatDateTime(i.DhEmi)
}

// Função para converter UFEnv para string
func (i *IdeOs) FormatUFEnv() {
	i.ProxyUFEnv = i.UFEnv.GetSigla()
}

// Função para converter UFIni para string
func (i *IdeOs) FormatUFIni() {
	i.ProxyUFIni = i.UFIni.GetSigla()
}

// Função para converter UFFim para string
func (i *IdeOs) FormatUFFim() {
	i.ProxyUFFim = i.UFFim.GetSigla()
}

// Função para formatar dhCont como string, se não for nulo
func (i *IdeOs) FormatDhCont() {
	if i.DhCont != nil {
		i.ProxyDhCont = utils.FormatDateTime(*i.DhCont)
	}
}
