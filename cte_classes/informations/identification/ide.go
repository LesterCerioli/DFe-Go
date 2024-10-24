package identification

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"time"

	"your_project_path/CTe/Classes/Informacoes/Tipos"
	"your_project_path/CTe/Classes/Servicos/Tipos"
	"your_project_path/DFe/Classes/Entidades"
	"your_project_path/DFe/Classes/Extensoes"
	"your_project_path/DFe/Classes/Flags"
)

// ide representa as informações de identificação do CT-e.
type ide struct {
	configuracaoServico *ConfiguracaoServico

	CUF                     Estado          `xml:"cUF"`
	CCT                     int             `xml:"-"`
	ProxyCCT                string          `xml:"cCT"`
	CFOP                    int             `xml:"CFOP"`
	NatOp                   string          `xml:"natOp"`
	ForPag                  *forPag         `xml:"-"`
	ForPagSpecified         bool            `xml:"-"`
	Mod                     ModeloDocumento `xml:"mod"`
	Serie                   int16           `xml:"serie"`
	NCT                     int64           `xml:"nCT"`
	DhEmi                   time.Time       `xml:"-"`
	ProxyDhEmi              string          `xml:"dhEmi"`
	TpImp                   tpImp           `xml:"tpImp"`
	TpEmis                  tpEmis          `xml:"tpEmis"`
	CDV                     int             `xml:"cDV"`
	TpAmb                   TipoAmbiente    `xml:"tpAmb"`
	TpCTe                   tpCTe           `xml:"tpCTe"`
	ProcEmi                 procEmi         `xml:"procEmi"`
	VerProc                 string          `xml:"verProc"`
	IndGlobalizado          *byte           `xml:"-"`
	IndGlobalizadoSpecified bool            `xml:"-"`
	RefCTE                  string          `xml:"refCTE"`
	CMunEnv                 int64           `xml:"cMunEnv"`
	XMunEnv                 string          `xml:"xMunEnv"`
	UFEnv                   Estado          `xml:"-"`
	ProxyUFEnv              string          `xml:"UFEnv"`
	Modal                   modal           `xml:"modal"`
	TpServ                  tpServ          `xml:"tpServ"`
	CMunIni                 int64           `xml:"cMunIni"`
	XMunIni                 string          `xml:"xMunIni"`
	UFIni                   Estado          `xml:"-"`
	ProxyUFIni              string          `xml:"UFIni"`
	CMunFim                 int64           `xml:"cMunFim"`
	XMunFim                 string          `xml:"xMunFim"`
	UFFim                   Estado          `xml:"-"`
	ProxyUFFim              string          `xml:"UFFim"`
	InfPercurso             []infPercurso   `xml:"infPercurso"`
	Retira                  retira          `xml:"retira"`
	XDetRetira              string          `xml:"xDetRetira"`
	IndIEToma               *indIEToma      `xml:"-"`
	IndIETomaSpecified      bool            `xml:"-"`
	TomaBase3               tomaBase3       `xml:"toma3"`
	Toma4                   toma4           `xml:"toma4"`
	DhCont                  *time.Time      `xml:"-"`
	ProxyDhCont             string          `xml:"dhCont"`
	XJust                   string          `xml:"xJust"`
}

// NewIde cria uma nova instância de ide com configuração de serviço opcional.
func NewIde(configuracaoServico *ConfiguracaoServico) *ide {
	if configuracaoServico == nil {
		configuracaoServico = ConfiguracaoServico.Instancia
	}
	return &ide{configuracaoServico: configuracaoServico}
}

// GetProxyCCT retorna o cCT formatado.
func (i *ide) GetProxyCCT() string {
	return fmt.Sprintf("%08d", i.CCT)
}

// SetProxyCCT define o cCT a partir do valor string.
func (i *ide) SetProxyCCT(value string) error {
	cct, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	i.CCT = cct
	return nil
}

// GetProxyDhEmi retorna a data de emissão formatada.
func (i *ide) GetProxyDhEmi() string {
	if i.configuracaoServico == nil {
		i.configuracaoServico = ConfiguracaoServico.Instancia
	}
	switch i.configuracaoServico.VersaoLayout {
	case versao.Ve200:
		return i.DhEmi.Format("2006-01-02T15:04:05") // Exemplo de formato para versão 2.0
	case versao.Ve400, versao.Ve300:
		return i.DhEmi.UTC().Format("2006-01-02T15:04:05Z07:00") // Exemplo de formato para versão 3.0 e 4.0
	default:
		return ""
	}
}

// SetProxyDhEmi define a data de emissão a partir do valor string.
func (i *ide) SetProxyDhEmi(value string) error {
	dhEmi, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return err
	}
	i.DhEmi = dhEmi
	return nil
}

// GetProxyDhCont retorna a data de contagem formatada.
func (i *ide) GetProxyDhCont() string {
	if i.DhCont == nil {
		return ""
	}
	return i.DhCont.UTC().Format("2006-01-02T15:04:05Z07:00")
}

// SetProxyDhCont define a data de contagem a partir do valor string.
func (i *ide) SetProxyDhCont(value string) error {
	if value == "" {
		i.DhCont = nil
		return nil
	}
	dhCont, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return err
	}
	i.DhCont = &dhCont
	return nil
}

// IndGlobalizadoSpecified verifica se indGlobalizado está especificado.
func (i *ide) IndGlobalizadoSpecified() bool {
	return i.IndGlobalizado != nil
}

// IndIETomaSpecified verifica se indIEToma está especificado.
func (i *ide) IndIETomaSpecified() bool {
	return i.IndIEToma != nil
}
