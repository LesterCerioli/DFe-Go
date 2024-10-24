package reception

import (
	"encoding/xml"
	"time"
)

// RetEnviCte representa a resposta do envio do CTe
type RetEnviCte struct {
	// AR02 - Versão do leiaute
	Versao string `xml:"versao,attr"` // Atributo

	// AR03 - Identificação do Ambiente: 1 – Produção / 2 - Homologação
	TpAmb TipoAmbiente `xml:"tpAmb"`

	// AR04 - Versão do Aplicativo que recebeu o Lote
	VerAplic string `xml:"verAplic"`

	// AR05 - Código do status da resposta
	CStat int `xml:"cStat"`

	// AR06 - Descrição literal do status da resposta
	DMotivo string `xml:"xMotivo"`

	// AR06a - Código da UF que atendeu a solicitação
	CUF Estado `xml:"cUF"`

	// AR06b - Data e Hora do Recebimento
	DhRecbto time.Time `xml:"-"`
	
	// ProxydhRecbto é um proxy para a conversão de DhRecbto em string
	ProxydhRecbto string `xml:"dhRecbto"`
	
	// AR07 - Dados do Recibo do Lote
	InfRec infRec `xml:"infRec"` // Certifique-se de que infRec esteja definido
}

// UnmarshalXML é usado para converter a string dhRecbto em um campo de tempo
func (r *RetEnviCte) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias RetEnviCte
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}

	// Converte o campo ProxydhRecbto para DhRecbto
	if aux.ProxydhRecbto != "" {
		dhRecbto, err := time.Parse("2006-01-02T15:04:05", aux.ProxydhRecbto)
		if err != nil {
			return err
		}
		r.DhRecbto = dhRecbto
	}

	return nil
}

// LoadXml desserializa uma string XML para um objeto RetEnviCte
func LoadXml(xmlData string) (*RetEnviCte, error) {
	var retorno RetEnviCte
	err := xml.Unmarshal([]byte(xmlData), &retorno)
	if err != nil {
		return nil, err
	}
	return &retorno, nil
}

// LoadXmlWithEnvio desserializa uma string XML e associa com o objeto enviCTe
func LoadXmlWithEnvio(xmlData string, enviCte *CTe) (*RetEnviCte, error) {
	retorno, err := LoadXml(xmlData)
	if err != nil {
		return nil, err
	}
	// Implementar a lógica para associar enviCte, se necessário
	return retorno, nil
}
