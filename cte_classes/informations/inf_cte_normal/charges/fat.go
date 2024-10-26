package charges

import (
	"fmt"
	"math"
)

// Fat representa informações sobre fatura no CT-e.
type Fat struct {
	nFat   string   // Número da fatura
	vOrig  *float64 // Valor original
	vDesc  *float64 // Valor do desconto
	vLiq   *float64 // Valor líquido
}

// SetVOrig define o valor original com arredondamento para 2 casas decimais.
func (f *Fat) SetVOrig(value float64) {
	rounded := round(value, 2)
	f.vOrig = &rounded
}

// VOrig retorna o valor original arredondado para 2 casas decimais.
func (f *Fat) VOrig() *float64 {
	if f.vOrig != nil {
		result := round(*f.vOrig, 2)
		return &result
	}
	return nil
}

// SetVDesc define o valor do desconto com arredondamento para 2 casas decimais.
func (f *Fat) SetVDesc(value float64) {
	rounded := round(value, 2)
	f.vDesc = &rounded
}

// VDesc retorna o valor do desconto arredondado para 2 casas decimais.
func (f *Fat) VDesc() *float64 {
	if f.vDesc != nil {
		result := round(*f.vDesc, 2)
		return &result
	}
	return nil
}

// SetVLiq define o valor líquido com arredondamento para 2 casas decimais.
func (f *Fat) SetVLiq(value float64) {
	rounded := round(value, 2)
	f.vLiq = &rounded
}

// VLiq retorna o valor líquido arredondado para 2 casas decimais.
func (f *Fat) VLiq() *float64 {
	if f.vLiq != nil {
		result := round(*f.vLiq, 2)
		return &result
	}
	return nil
}

// VOrigSpecified indica se o valor original está definido.
func (f *Fat) VOrigSpecified() bool {
	return f.vOrig != nil
}

// VDescSpecified indica se o valor do desconto está definido.
func (f *Fat) VDescSpecified() bool {
	return f.vDesc != nil
}

// VLiqSpecified indica se o valor líquido está definido.
func (f *Fat) VLiqSpecified() bool {
	return f.vLiq != nil
}

// round arredonda um número para o número de casas decimais especificado.
func round(val float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(val*multiplier) / multiplier
}

// ToString imprime a representação da fatura.
func (f *Fat) ToString() string {
	return fmt.Sprintf("nFat: %s, vOrig: %.2f, vDesc: %.2f, vLiq: %.2f",
		f.nFat, f.VOrig(), f.VDesc(), f.VLiq())
}
