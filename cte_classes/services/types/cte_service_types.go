package types

import (
	"encoding/xml"
	"fmt"
)

// ServicoCTe representa os diferentes tipos de serviços do CTe.
type ServicoCTe int

const (
	RecepcaoEventoCancelmento ServicoCTe = iota
	RecepcaoEventoCartaCorrecao
	RecepcaoEventoEpec
	RecepcaoEventoManifestacaoDestinatario
	CteRecepcao
	CteRetRecepcao
	CteConsultaCadastro
	CteInutilizacao
	CteConsultaProtocolo
	CteStatusServico
	CteAutorizacao
	CteRetAutorizacao
	CteDistribuicaoDFe
	CteConsultaDest
	CteDownloadNF
)

// TipoRecepcaoEvento representa os tipos de recepção de evento.
type TipoRecepcaoEvento int

const (
	Nenhum TipoRecepcaoEvento = iota
	Cancelamento
	CartaCorrecao
	Epec
	ManifestacaoDestinatario
)

// TipoManifestacao representa os diferentes tipos de manifestação.
type TipoManifestacao int

const (
	e210200 TipoManifestacao = iota + 210200
	e210210
	e210220
	e210240
)

// Versao representa as versões do leiaute.
type Versao int

const (
	ve200 Versao = iota + 200
	ve300
	ve400
)

// IndicadorSincronizacao representa o indicador de sincronização.
type IndicadorSincronizacao int

const (
	Assincrono IndicadorSincronizacao = iota
	Sincrono
)

// Description é um auxiliar para gerar descrições para os enums.
func (s ServicoCTe) Description() string {
	switch s {
	case RecepcaoEventoCancelmento:
		return "Serviço destinado à recepção de mensagem do Evento de Cancelamento da NF-e"
	case RecepcaoEventoCartaCorrecao:
		return "Serviço destinado à recepção de mensagem do Evento de Carta de Correção da NF-e"
	case RecepcaoEventoEpec:
		return "Serviço destinado à recepção de mensagem do Evento EPEC da NF-e"
	case RecepcaoEventoManifestacaoDestinatario:
		return "Serviço destinado à recepção de mensagem do Evento de Manifestação do destinatário da NF-e"
	case CteRecepcao:
		return "Serviço destinado à recepção de mensagens de lote de NF-e versão 2.0"
	case CteRetRecepcao:
		return "Serviço destinado a retornar o resultado do processamento do lote de NF-e versão 2.0"
	case CteConsultaCadastro:
		return "Serviço para consultar o cadastro de contribuintes do ICMS da unidade federada"
	case CteInutilizacao:
		return "Serviço destinado ao atendimento de solicitações de inutilização de numeração"
	case CteConsultaProtocolo:
		return "Serviço destinado ao atendimento de solicitações de consulta da situação atual da NF-e na Base de Dados do Portal da Secretaria de Fazenda Estadual"
	case CteStatusServico:
		return "Serviço destinado à consulta do status do serviço prestado pelo Portal da Secretaria de Fazenda Estadual"
	case CteAutorizacao:
		return "Serviço destinado à recepção de mensagens de lote de NF-e versão 3.10"
	case CteRetAutorizacao:
		return "Serviço destinado a retornar o resultado do processamento do lote de NF-e versão 3.10"
	case CteDistribuicaoDFe:
		return "Distribui documentos e informações de interesse do ator da NF-e"
	case CteConsultaDest:
		return "Serviço de Consulta da Relação de Documentos Destinados para um determinado CNPJ de destinatário informado na NF-e"
	case CteDownloadNF:
		return "Serviço destinado ao atendimento de solicitações de download de Notas Fiscais Eletrônicas por seus destinatários"
	default:
		return "Serviço desconhecido"
	}
}

// XML é uma estrutura que contém informações para serialização XML.
type XML struct {
	XMLName xml.Name `xml:"ServicoCTe"`
}

func main() {
	// Exemplo de uso
	servico := RecepcaoEventoCancelmento
	fmt.Println(servico.Description())
}
