package event

import (
	"encoding/xml"
	"errors"
	"time"
)

// InfEventoEnv representa o grupo de informações do registro do Evento
type InfEventoEnv struct {
	ConfiguracaoServico *ConfiguracaoServico `xml:"-"` // Configuração do serviço
	ID                   string                `xml:"Id,attr"`                      // HP07 - ID do evento
	COrgao               Estado                `xml:"cOrgao"`                       // HP08 - Código do órgão de recepção do Evento
	TpAmb               TipoAmbiente          `xml:"tpAmb"`                        // HP09 - Identificação do Ambiente
	CNPJ                 string                `xml:"CNPJ"`                         // HP10 - CNPJ do autor do Evento
	CPF                  string                `xml:"CPF"`                          // CPF do autor do Evento
	ChCTe                string                `xml:"chCTe"`                        // HP12 - Chave de Acesso da NF-e vinculada ao Evento
	DhEvento             time.Time             `xml:"dhEvento"`                     // HP13 - Data e hora do evento
	TpEvento             CTeTipoEvento         `xml:"tpEvento"`                     // HP14 - Código do evento
	NSeqEvento           int                   `xml:"nSeqEvento"`                   // HP15 - Sequencial do evento
	DetEvento            *DetEvento            `xml:"detEvento"`                    // HP17 - Informações do Pedido de Cancelamento
}

// NewInfEventoEnv cria uma nova instância de InfEventoEnv com configuração do serviço
func NewInfEventoEnv(configuracaoServico *ConfiguracaoServico) *InfEventoEnv {
	if configuracaoServico == nil {
		configuracaoServico = ConfiguracaoServico.Instancia
	}
	return &InfEventoEnv{ConfiguracaoServico: configuracaoServico}
}

// ProxyDhEvento converte dhEvento para string de acordo com a versão do layout
func (e *InfEventoEnv) ProxyDhEvento() (string, error) {
	if e.ConfiguracaoServico == nil {
		e.ConfiguracaoServico = ConfiguracaoServico.Instancia
	}
	switch e.ConfiguracaoServico.VersaoLayout {
	case VersaoVe200:
		return e.DhEvento.Format("2006-01-02 15:04:05"), nil // Exemplo de formatação sem UTC
	case VersaoVe400, VersaoVe300:
		return e.DhEvento.Format("2006-01-02T15:04:05Z07:00"), nil // Formato UTC
	default:
		return "", errors.New("versão inválida para CT-e")
	}
}
