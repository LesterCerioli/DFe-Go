package inf_modals

import (
	"time"
)

// Rodo representa as informações do transporte rodoviário.
type Rodo struct {
	RNTRC   string     `xml:"RNTRC"`
	DPrev   *time.Time `xml:"dPrev,omitempty"`
	Lota    *Lota      `xml:"lota,omitempty"`
	CIOT    string     `xml:"CIOT"`
	Occ     []Occ      `xml:"occ"`
	ValePed []ValePed  `xml:"valePed"`
	Veic    []Veic     `xml:"veic"`
	LacRodo []LacRodo  `xml:"lacRodo"`
	Moto    []Moto     `xml:"moto"`
}

// ProxydPrev retorna a representação em string da data prevista.
func (r *Rodo) ProxydPrev() string {
	if r.DPrev == nil {
		return ""
	}
	return r.DPrev.Format("2006-01-02") // Formato de data YYYY-MM-DD
}

// ShouldSerializeVeic verifica se deve serializar o campo veic.
func (r *Rodo) ShouldSerializeVeic() bool {
	return r.Veic != nil
}

// ShouldSerializeMoto verifica se deve serializar o campo moto.
func (r *Rodo) ShouldSerializeMoto() bool {
	return r.Moto != nil
}

// Lota, Occ, ValePed, Veic, LacRodo e Moto devem ser definidos como structs separadas
// de acordo com suas definições em C#.
type Lota struct {
	// Definições do tipo Lota
}

type Occ struct {
	// Definições do tipo Occ
}

type ValePed struct {
	// Definições do tipo ValePed
}

type Veic struct {
	// Definições do tipo Veic
}

type LacRodo struct {
	NLacre string `xml:"nLacre"`
}

type Moto struct {
	XNome string `xml:"xNome"`
	CPF   string `xml:"CPF"`
}
