package inf_cte_normal

import (
	"time"
)

// InfDocRef representa informações de referência de documentos.
type InfDocRef struct {
	NDoc              string    `xml:"nDoc"`     // Número do Documento
	Serie             int16     `xml:"serie"`    // Série do Documento
	Subserie          *int16    `xml:"subserie"` // Sub-série do Documento (opcional)
	SubserieSpecified bool      // Indica se a sub-série está especificada
	DEmi              time.Time `xml:"-"`
	ProxydEmi         string    `xml:"dEmi"` // Data de Emissão como string
	VDoc              *float64  // Valor do Documento (opcional)
	VDocSpecified     bool      // Indica se o valor do documento está especificado
}

// ParaDataString é uma função auxiliar para formatar a data como string.
func (d time.Time) ParaDataString() string {
	// Implemente a lógica de formatação da data conforme necessário.
	return d.Format("2006-01-02") // Exemplo de formatação
}

// Arredondar é uma função auxiliar para arredondar um valor decimal.
func Arredondar(value float64, precision int) float64 {
	// Implemente a lógica de arredondamento conforme necessário.
	// Exemplo: arredondar para 2 casas decimais.
	return math.Round(value*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
}
