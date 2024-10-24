package charges

import (
	"encoding/xml"
	"fmt"
	"math"
	"time"
)

// dup representa uma duplicata no CT-e.
type Dup struct {
	NDup      string    `xml:"nDup"`      // Número da duplicata
	DVenc     *time.Time `xml:"-"`        // Data de vencimento
	vDup      *float64  `xml:"-"`        // Valor da duplicata
}

// ProxydVenc retorna a data de vencimento como uma string formatada.
func (d *Dup) ProxydVenc() string {
	if d.DVenc == nil {
		return ""
	}
	return d.DVenc.Format("2006-01-02") // Formato de data padrão
}

// SetDVenc define a data de vencimento a partir de uma string.
func (d *Dup) SetDVenc(value string) error {
	t, err := time.Parse("2006-01-02", value) // Formato de data esperado
	if err != nil {
		return err
	}
	d.DVenc = &t
	return nil
}

// VDup retorna o valor da duplicata arredondado para 2 casas decimais.
func (d *Dup) VDup() *float64 {
	if d.vDup != nil {
		val := *d.vDup
		result := round(val, 2)
		return &result
	}
	return nil
}

// SetVDup define o valor da duplicata arredondado para 2 casas decimais.
func (d *Dup) SetVDup(value float64) {
	rounded := round(value, 2)
	d.vDup = &rounded
}

// VDupSpecified indica se o valor da duplicata está definido.
func (d *Dup) VDupSpecified() bool {
	return d.vDup != nil
}

// round arredonda um número para o número de casas decimais especificado.
func round(val float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(val*multiplier) / multiplier
}

// ToString imprime a representação da duplicata.
func (d *Dup) ToString() string {
	return fmt.Sprintf("nDup: %s, dVenc: %s, vDup: %.2f", d.NDup, d.ProxydVenc(), d.VDup())
}
