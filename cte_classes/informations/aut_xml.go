package informations

import (
	"errors"
)

// AutXML representa um objeto que contém informações de CNPJ e CPF.
type AutXML struct {
	cnpj string
	cpf  string
}

// ErroCpfCnpjPreenchidos é a mensagem de erro quando ambos CNPJ e CPF são preenchidos.
const ErroCpfCnpjPreenchidos = "Somente preencher um dos campos: CNPJ ou CPF, para um objeto do tipo AutXML!"

// SetCNPJ define o valor do CNPJ, garantindo que o CPF não esteja preenchido.
func (a *AutXML) SetCNPJ(cnpj string) error {
	if cnpj == "" {
		return nil
	}
	if a.cpf == "" {
		a.cnpj = cnpj
		return nil
	}
	return errors.New(ErroCpfCnpjPreenchidos)
}

// GetCNPJ retorna o valor do CNPJ.
func (a *AutXML) GetCNPJ() string {
	return a.cnpj
}

// SetCPF define o valor do CPF, garantindo que o CNPJ não esteja preenchido.
func (a *AutXML) SetCPF(cpf string) error {
	if cpf == "" {
		return nil
	}
	if a.cnpj == "" {
		a.cpf = cpf
		return nil
	}
	return errors.New(ErroCpfCnpjPreenchidos)
}

// GetCPF retorna o valor do CPF.
func (a *AutXML) GetCPF() string {
	return a.cpf
}
