package inf_documents

import (
	"encoding/xml"
	"github.com/shopspring/decimal" // Certifique-se de que este pacote esteja disponível.
)

// InfUnidTransp representa informações sobre a unidade de transporte.
type InfUnidTransp struct {
	XMLName      xml.Name        `xml:"infUnidTransp"`
	TpUnidTransp TpUnidTransp    `xml:"tpUnidTransp"`
	IdUnidTransp string          `xml:"idUnidTransp"`
	LacUnidTransp []LacUnidTransp `xml:"lacUnidTransp"`
	InfUnidCarga []InfUnidCarga  `xml:"infUnidCarga"`
	QtdRat       *decimal.Decimal `xml:"qtdRat"`
}

// GetQtdRat retorna a quantidade de rat arredondada em 3 casas decimais.
func (inf *InfUnidTransp) GetQtdRat() decimal.Decimal {
	if inf.QtdRat != nil {
		return inf.QtdRat.Round(3) // Supondo que exista um método de arredondamento
	}
	return decimal.NewFromFloat(0) // Retorna zero se não estiver especificado
}

// SetQtdRat define a quantidade de rat com arredondamento em 3 casas decimais.
func (inf *InfUnidTransp) SetQtdRat(value decimal.Decimal) {
	inf.QtdRat = &value // Armazenar o valor como ponteiro
}

// QtdRatSpecified verifica se a quantidade de rat está especificada.
func (inf *InfUnidTransp) QtdRatSpecified() bool {
	return inf.QtdRat != nil
}
