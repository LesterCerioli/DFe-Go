package event

import (
	"encoding/xml"
	"time"
)

// InfEventoRet representa a resposta de um evento
type InfEventoRet struct {
	ID          string            `xml:"Id,attr"`                          // HR12 - Identificador da TAG a ser assinada
	TpAmb      TipoAmbiente      `xml:"tpAmb"`                            // HR13 - Identificação do Ambiente
	VerAplic    string            `xml:"verAplic"`                         // HR14 - Versão da aplicação que registrou o Evento
	COrgao      Estado            `xml:"cOrgao"`                           // HR15 - Código da UF que registrou o Evento
	CStat       int               `xml:"cStat"`                            // HR16 - Código do status da resposta
	 XMotivo    string            `xml:"xMotivo"`                          // HR17 - Descrição do status da resposta
	ChCTe       string            `xml:"chCTe"`                            // HR18 - Chave de Acesso da NF-e vinculada ao evento
	TpEvento    *int              `xml:"tpEvento"`                         // HR19 - Código do Tipo do Evento
	XEvento     string            `xml:"xEvento"`                          // HR20 - Descrição do Evento
	NSeqEvento  *int              `xml:"nSeqEvento"`                       // HR21 - Sequencial do evento
	DhRegEvento *time.Time        `xml:"dhRegEvento"`                      // HR25 - Data e hora de registro do evento
	NProt       string            `xml:"nProt"`                            // HR26 - Número do Protocolo da NF-e
	Signature   *Signature        `xml:"Signature"`                       // HR27 - Assinatura Digital do documento XML
}

// ProxyDhRegEvento converte DhRegEvento para string no formato UTC
func (e *InfEventoRet) ProxyDhRegEvento() string {
	if e.DhRegEvento != nil {
		return e.DhRegEvento.Format(time.RFC3339) // Formato UTC
	}
	return ""
}

// ShouldSerializeTpEvento verifica se tpEvento deve ser serializado
func (e *InfEventoRet) ShouldSerializeTpEvento() bool {
	return e.TpEvento != nil
}

// ShouldSerializeNSeqEvento verifica se nSeqEvento deve ser serializado
func (e *InfEventoRet) ShouldSerializeNSeqEvento() bool {
	return e.NSeqEvento != nil
}
