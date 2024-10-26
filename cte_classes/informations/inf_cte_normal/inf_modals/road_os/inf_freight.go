package road_os

import (
	"time"
)

// TpFretamento é um tipo definido para representar o tipo de fretamento.
type TpFretamento int

// InfFretamento representa as informações de fretamento.
type InfFretamento struct {
	TpFretamento TpFretamento `json:"tpFretamento" xml:"tpFretamento"`
	dhViagem     *time.Time   `json:"-" xml:"-"`
}

// ProxyDHViagem retorna a data e hora da viagem no formato de string.
func (i *InfFretamento) ProxyDHViagem() string {
	if i.dhViagem != nil {
		return i.dhViagem.UTC().Format(time.RFC3339) // Formato UTC
	}
	return ""
}

// SetProxyDHViagem converte a string para o tipo time.Time.
func (i *InfFretamento) SetProxyDHViagem(value string) {
	if value == "" {
		i.dhViagem = nil
		return
	}
	t, err := time.Parse(time.RFC3339, value) // Altere o formato conforme necessário
	if err == nil {
		i.dhViagem = &t
	}
}
