package inf_cte_subs

import (
	"time"

	"CTe/Classes/Informacoes/Tipos"
)

// RefNF representa a referência de nota fiscal.
type RefNF struct {
	CNPJ        string  `json:"CNPJ"`         // CNPJ do emitente
	CPF         string  `json:"CPF"`          // CPF do emitente
	Mod         string  `json:"mod"`          // Modelos da nota fiscal
	Serie       int16   `json:"serie"`        // Série da nota fiscal
	Subserie    *int16  `json:"subserie"`     // Subserie da nota fiscal, pode ser nulo
	Nro         int     `json:"nro"`          // Número da nota fiscal
	valor       float64 `json:"valor"`        // Valor da nota fiscal
	dEmi        time.Time // Data de emissão (não serializado diretamente)
}

// SetValor define o valor da nota fiscal com arredondamento
func (ref *RefNF) SetValor(value float64) {
	ref.valor = arredondar(value, 2)
}

// GetValor retorna o valor arredondado da nota fiscal
func (ref *RefNF) GetValor() float64 {
	return arredondar(ref.valor, 2)
}

// SetDEmi define a data de emissão a partir de uma string
func (ref *RefNF) SetDEmi(value string) error {
	parsedDate, err := time.Parse("2006-01-02", value) // Altere o layout conforme necessário
	if err != nil {
		return err
	}
	ref.dEmi = parsedDate
}

// GetDEmi retorna a data de emissão como string
func (ref *RefNF) GetDEmi() string {
	return ref.dEmi.Format("2006-01-02") // Altere o formato conforme necessário
}

// arredondar é uma função auxiliar para arredondar valores para um número específico de casas decimais
func arredondar(val float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	return math.Round(val*pow) / pow
}
