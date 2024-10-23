package inf_documents

import (
	"encoding/xml"
	"time"
)

// InfNF representa as informações da Nota Fiscal.
type InfNF struct {
	XMLName       xml.Name        `xml:"infNF"`
	NRoma        string           `xml:"nRoma"`
	NPed         string           `xml:"nPed"`
	Mod          Mod              `xml:"mod"`
	Serie        string           `xml:"serie"`
	NDoc         string           `xml:"nDoc"`
	DEmi         time.Time        `xml:"-"`
	ProxydEmi    string           `xml:"dEmi"` // Formatação deve ser tratada ao serializar

	VBC          decimal          `xml:"vBC"`
	VICMS        decimal          `xml:"vICMS"`
	VBCST        decimal          `xml:"vBCST"`
	VST          decimal          `xml:"vST"`
	VProd        decimal          `xml:"vProd"`
	VNF          decimal          `xml:"vNF"`
	NCFOP        int              `xml:"nCFOP"`
	NPeso        *decimal         `xml:"nPeso,omitempty"`
	NPesoSpecified bool           `xml:"-"`
	PIN          string           `xml:"PIN"`
	DPrev        *time.Time       `xml:"-"`
	ProxydPrev   string           `xml:"dPrev"` // Formatação deve ser tratada ao serializar

	InfUnidTransp []InfUnidTransp `xml:"infUnidTransp"`
	InfUnidCarga  []InfUnidCarga   `xml:"infUnidCarga"`
}

// Decimal é um tipo personalizado para representar valores decimais.
type Decimal float64

// Arredondar arredonda o valor decimal para o número de casas decimais especificadas.
func (d Decimal) Arredondar(casas int) Decimal {
	// Implementação da função de arredondamento (considerar a precisão desejada)
	return d
}
