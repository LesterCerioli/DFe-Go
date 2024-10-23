package inf_resp_tec

import (
	"encoding/xml"
	"fmt"
)

// InfRespTec representa as informações do responsável técnico.
type InfRespTec struct {
	CNPJ      string  `xml:"CNPJ"`
	XContato  string  `xml:"xContato"`
	Email     string  `xml:"email"`
	Fone      string  `xml:"fone"`
	IdCSRT    *int    `xml:"idCSRT,omitempty"` // Campo opcional que pode ser nulo
	HashCSRT  string  `xml:"hashCSRT"`
}

// IdCSRTSpecified verifica se IdCSRT possui um valor.
func (i *InfRespTec) IdCSRTSpecified() bool {
	return i.IdCSRT != nil
}

// ProxyIdCSRT obtém o valor de IdCSRT formatado como string.
func (i *InfRespTec) ProxyIdCSRT() string {
	if i.IdCSRT != nil {
		return fmt.Sprintf("%03d", *i.IdCSRT) // Formata para três dígitos
	}
	return ""
}

// SetProxyIdCSRT define o valor de IdCSRT a partir de uma string.
func (i *InfRespTec) SetProxyIdCSRT(value string) error {
	var id int
	_, err := fmt.Sscanf(value, "%d", &id) // Converte a string para int
	if err != nil {
		return err // Retorna erro se a conversão falhar
	}
	i.IdCSRT = &id // Atribui o valor a IdCSRT
	return nil
}
