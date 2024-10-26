package inf_documents

import (
	"encoding/xml"
	"time"
)

// InfNFe representa as informações da Nota Fiscal Eletrônica.
type InfNFe struct {
	XMLName        xml.Name        `xml:"infNFe"`
	Chave          string          `xml:"chave"`
	PIN            string          `xml:"PIN"`
	DPrev          *time.Time      `xml:"-"`
	ProxydPrev     string          `xml:"dPrev"` // Formatação deve ser tratada ao serializar
	InfUnidTransp  []InfUnidTransp `xml:"infUnidTransp"`
	InfUnidCarga   []InfUnidCarga  `xml:"infUnidCarga"`
}

// FormatarData retorna a data em formato string no padrão AAAA-MM-DD.
func (inf *InfNFe) FormatarData() string {
	if inf.DPrev != nil {
		return inf.DPrev.Format("2006-01-02") // Formato YYYY-MM-DD
	}
	return ""
}

// ParseData configura dPrev a partir de uma string.
func (inf *InfNFe) ParseData(data string) error {
	parsedDate, err := time.Parse("2006-01-02", data)
	if err != nil {
		return err
	}
	inf.DPrev = &parsedDate
	return nil
}
