package previous_docs

import (
	"time"
)

// IDDocAntPap representa informações sobre documentos anteriores em papel.
type IDDocAntPap struct {
	TpDoc   TpDocAnterior `xml:"tpDoc"`   // Tipo do documento anterior
	Serie   string        `xml:"serie"`   // Série do documento
	Subser  string        `xml:"subser"`   // Sub-série do documento
	NDoc    string        `xml:"nDoc"`    // Número do documento
	Demi    *time.Time    `xml:"dEmi"`    // Data de emissão do documento (pointer para permitir nil)
}

// ProxydEmi retorna a data de emissão como string formatada
func (doc *IDDocAntPap) ProxydEmi() string {
	if doc.Demi != nil {
		return doc.Demi.Format("2006-01-02") // formato YYYY-MM-DD, ajuste conforme necessário
	}
	return ""
}

// SetProxydEmi define a data de emissão a partir de uma string
func (doc *IDDocAntPap) SetProxydEmi(value string) error {
	if value == "" {
		doc.Demi = nil
		return nil
	}
	date, err := time.Parse("2006-01-02", value) // ajuste o formato conforme necessário
	if err != nil {
		return err
	}
	doc.Demi = &date
	return nil
}
