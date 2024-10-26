package protocols

import (
	"time"
)

// InfProt representa a estrutura das informações do protocolo.
type InfProt struct {
	// PR04 - Identificador da TAG a ser assinada, somente precisa ser informado se a UF assinar a resposta.
	Id string `json:"id" xml:"id"`

	// PR05 - Identificação do Ambiente
	TpAmb TipoAmbiente `json:"tpAmb" xml:"tpAmb"`

	// PR06 - Versão do Aplicativo que processou a consulta.
	VerAplic string `json:"verAplic" xml:"verAplic"`

	// PR07 - Chave de Acesso da CT-e
	ChCTe string `json:"chCTe" xml:"chCTe"`

	// PR08 - Data e hora de recebimento
	DhRecbto time.Time `json:"-" xml:"-"`

	// Proxy para a data de recebimento em formato string
	ProxydhRecbto string `json:"dhRecbto" xml:"dhRecbto"`

	// PR09 - Número do Protocolo da CT-e
	NProt string `json:"nProt" xml:"nProt"`

	// PR10 - Digest Value da CT-e processada
	DigVal string `json:"digVal" xml:"digVal"`

	// PR11 - Código do status da resposta.
	CStat int `json:"cStat" xml:"cStat"`

	// PR12 - Descrição literal do status da resposta.
	XMotivo string `json:"xMotivo" xml:"xMotivo"`

	InfFisco InfFisco `json:"infFisco" xml:"infFisco"`

	// PR13 - Assinatura XML do grupo identificado pelo atributo “Id”
	Signature Signature `json:"signature" xml:"signature"`
}

// Método para converter ProxydhRecbto para dhRecbto
func (i *InfProt) SetProxydhRecbto(value string) {
	i.DhRecbto, _ = time.Parse(time.RFC3339, value) // Certifique-se de usar o formato adequado
}

func (i *InfProt) GetProxydhRecbto() string {
	return i.DhRecbto.Format(time.RFC3339)
}
