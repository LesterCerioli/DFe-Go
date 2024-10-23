package inf_documents

import (
	"encoding/xml"
)

// InfUnidCarga representa informações sobre a unidade de carga.
type InfUnidCarga struct {
	XMLName      xml.Name        `xml:"infUnidCarga"`
	TpUnidCarga  TpUnidCarga     `xml:"tpUnidCarga"`
	IdUnidCarga  string          `xml:"idUnidCarga"`
	LacUnidCarga []LacUnidCarga  `xml:"lacUnidCarga"`
	QtdRat       *decimal.Decimal `xml:"qtdRat"`
}

// GetQtdRat retorna a quantidade de rat arredondada em 3 casas decimais.
func (inf *InfUnidCarga) GetQtdRat() decimal.Decimal {
	if inf.QtdRat != nil {
		return inf.QtdRat.Arredondar(3)
	}
	return 0
}

// SetQtdRat define a quantidade de rat com arredondamento em 3 casas decimais.
func (inf *InfUnidCarga) SetQtdRat(value decimal.Decimal) {
	inf.QtdRat = value.Arredondar(3)
}

// QtdRatSpecified verifica se a quantidade de rat está especificada.
func (inf *InfUnidCarga) QtdRatSpecified() bool {
	return inf.QtdRat != nil
}
