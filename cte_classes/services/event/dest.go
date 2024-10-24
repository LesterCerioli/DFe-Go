package event

import (
	"errors"
	"fmt"
	"math"
)

// Erros relacionados ao preenchimento de CNPJ, CPF ou idEstrangeiro
const ErroCpfCnpjIdEstrangeiroPreenchidos = "Somente preencher um dos campos: CNPJ, CPF ou idEstrangeiro, para um objeto do tipo dest!"

// Dest é a struct que representa o destinatário de um evento.
type Dest struct {
	UF            string  // Sigla da UF do destinatário (informar "EX" para operações exteriores)
	CNPJ          string  // CNPJ do destinatário
	CPF           string  // CPF do destinatário
	IdEstrangeiro string  // Identificação estrangeira do destinatário
	IE            string  // Inscrição Estadual do destinatário (omitido se "ISENTO")
	VNF           float64 // Valor total da NF-e
	VICMS         float64 // Valor total do ICMS
	VST           float64 // Valor total do ICMS de Substituição Tributária
}

// SetCNPJ define o CNPJ, garantindo que CPF e idEstrangeiro estejam vazios.
func (d *Dest) SetCNPJ(cnpj string) error {
	if cnpj != "" {
		if d.CPF == "" && d.IdEstrangeiro == "" {
			d.CNPJ = cnpj
		} else {
			return errors.New(ErroCpfCnpjIdEstrangeiroPreenchidos)
		}
	}
	return nil
}

// SetCPF define o CPF, garantindo que CNPJ e idEstrangeiro estejam vazios.
func (d *Dest) SetCPF(cpf string) error {
	if cpf != "" {
		if d.CNPJ == "" && d.IdEstrangeiro == "" {
			d.CPF = cpf
		} else {
			return errors.New(ErroCpfCnpjIdEstrangeiroPreenchidos)
		}
	}
	return nil
}

// SetIdEstrangeiro define o IdEstrangeiro, garantindo que CNPJ e CPF estejam vazios.
func (d *Dest) SetIdEstrangeiro(id string) error {
	if id != "" {
		if d.CNPJ == "" && d.CPF == "" {
			d.IdEstrangeiro = id
		} else {
			return errors.New(ErroCpfCnpjIdEstrangeiroPreenchidos)
		}
	}
	return nil
}

// SetVNF define o valor total da NF-e arredondado com duas casas decimais.
func (d *Dest) SetVNF(vnf float64) {
	d.VNF = round(vnf, 2)
}

// SetVICMS define o valor total do ICMS arredondado com duas casas decimais.
func (d *Dest) SetVICMS(vicms float64) {
	d.VICMS = round(vicms, 2)
}

// SetVST define o valor total do ICMS de Substituição Tributária arredondado com duas casas decimais.
func (d *Dest) SetVST(vst float64) {
	d.VST = round(vst, 2)
}

// ShouldSerializeIE verifica se a Inscrição Estadual deve ser serializada.
func (d *Dest) ShouldSerializeIE() bool {
	return d.IE != "" && d.IE != "ISENTO"
}

// Função utilitária para arredondar valores decimais com precisão definida.
func round(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
