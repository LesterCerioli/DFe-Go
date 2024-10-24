package inf_documents

import (
	"encoding/xml"
	"time"
)

// InfOutros representa informações de documentos adicionais.
type InfOutros struct {
	XMLName       xml.Name        `xml:"infOutros"`
	TpDoc         TpDoc           `xml:"tpDoc"`
	DescOutros    string          `xml:"descOutros"`
	NDoc          string          `xml:"nDoc"`
	DEmi          *time.Time      `xml:"-"`
	ProxydEmi     string          `xml:"dEmi"` // Formatação deve ser tratada ao serializar
	VDocFisc      *decimal.Decimal `xml:"vDocFisc"`
	VDocFiscSpecified bool         `xml:"-"`
	DPrev         *time.Time      `xml:"-"`
	ProxydPrev    string          `xml:"dPrev"` // Formatação deve ser tratada ao serializar
	InfUnidTransp []InfUnidTransp `xml:"infUnidTransp"`
	InfUnidCarga  []InfUnidCarga  `xml:"infUnidCarga"`
}

// FormatarData retorna a data em formato string no padrão AAAA-MM-DD.
func (inf *InfOutros) FormatarDataDemi() string {
	if inf.DEmi != nil {
		return inf.DEmi.Format("2006-01-02") // Formato YYYY-MM-DD
	}
	return ""
}

// ParseDataDemi configura DEmi a partir de uma string.
func (inf *InfOutros) ParseDataDemi(data string) error {
	parsedDate, err := time.Parse("2006-01-02", data)
	if err != nil {
		return err
	}
	inf.DEmi = &parsedDate
	return nil
}

// FormatarDataDPrev retorna a data em formato string no padrão AAAA-MM-DD.
func (inf *InfOutros) FormatarDataDPrev() string {
	if inf.DPrev != nil {
		return inf.DPrev.Format("2006-01-02") // Formato YYYY-MM-DD
	}
	return ""
}

// ParseDataDPrev configura DPrev a partir de uma string.
func (inf *InfOutros) ParseDataDPrev(data string) error {
	parsedDate, err := time.Parse("2006-01-02", data)
	if err != nil {
		return err
	}
	inf.DPrev = &parsedDate
	return nil
}
