package inf_cte_anu

import (
	"time"
)

// InfCteAnu representa a estrutura das informações de cancelamento do CT-e.
type InfCteAnu struct {
	ChCte    string    // Chave do CT-e
	dEmi     time.Time // Data de emissão
	ProxydEmi string   // ProxydEmi representa a data de emissão como string
}

// SetProxydEmi define a data de emissão a partir de uma string.
func (i *InfCteAnu) SetProxydEmi(value string) error {
	// Tenta converter a string para time.Time e atribui a dEmi.
	dEmi, err := time.Parse("2006-01-02", value) // Formato de data "YYYY-MM-DD"
	if err != nil {
		return err // Retorna erro se a conversão falhar
	}
	i.dEmi = dEmi
	return nil
}

// GetProxydEmi obtém a data de emissão como string.
func (i *InfCteAnu) GetProxydEmi() string {
	return i.dEmi.Format("2006-01-02") // Formata a data para string
}
