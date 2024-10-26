package reception

import (
	"encoding/xml"
	"time"
)

// InfRec representa a informação de recibo do CTe
type InfRec struct {
	// AR08 - Número do Recibo gerado pelo Portal da Secretaria de Fazenda Estadual
	NRec string `xml:"nRec"` // O número do recibo

	// AR10 - Tempo médio de resposta do serviço (em segundos) dos últimos 5 minutos
	TMed int `xml:"tMed"` // O tempo médio de resposta

	// Data e hora do recebimento
	DhRecbto time.Time `xml:"-"` // Ignorado na serialização

	// Proxy para a propriedade DhRecbto, usado para serialização
	ProxydhRecbto string `xml:"dhRecbto"`
}

// UnmarshalXML trata a deserialização de dhRecbto
func (ir *InfRec) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias InfRec // Cria um alias para evitar recursão
	if err := d.DecodeElement((*Alias)(ir), &start); err != nil {
		return err
	}
	// Converte a string ProxydhRecbto para time.Time
	if ir.ProxydhRecbto != "" {
		parsedTime, err := time.Parse(time.RFC3339, ir.ProxydhRecbto)
		if err != nil {
			return err
		}
		ir.DhRecbto = parsedTime
	}
	return nil
}

// MarshalXML trata a serialização de dhRecbto
func (ir *InfRec) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	ir.ProxydhRecbto = ir.DhRecbto.Format(time.RFC3339) // Formato UTC
	return e.EncodeElement(ir, start)
}
