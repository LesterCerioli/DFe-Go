package inf_modals

import (
	"fmt"
	"math"
	"time"
)

// Duto representa o modal duto.
type Duto struct {
	VTar      *float64  `xml:"vTar"`
	DIni      time.Time `xml:"-"`
	ProxydIni string    `xml:"dIni"`
	DFim      time.Time `xml:"-"`
	ProxydFim string    `xml:"dFim"`
}

// Arredondar arredonda o valor para o número de casas decimais especificadas.
func Arredondar(valor *float64, casasDecimais int) *float64 {
	if valor == nil {
		return nil
	}
	fator := math.Pow(10, float64(casasDecimais))
	novoValor := math.Round(*valor*fator) / fator
	return &novoValor
}

// GetVTar retorna o valor arredondado de vTar.
func (d *Duto) GetVTar() *float64 {
	return Arredondar(d.VTar, 6)
}

// SetVTar define o valor de vTar com arredondamento.
func (d *Duto) SetVTar(valor *float64) {
	d.VTar = Arredondar(valor, 6)
}

// ParaDataString converte time.Time para string no formato de data.
func ParaDataString(t time.Time) string {
	return t.Format("2006-01-02")
}

// ProxydIni retorna dIni como string.
func (d *Duto) GetProxydIni() string {
	return ParaDataString(d.DIni)
}

// SetProxydIni define dIni a partir de uma string.
func (d *Duto) SetProxydIni(valor string) {
	d.DIni, _ = time.Parse("2006-01-02", valor)
}

// ProxydFim retorna dFim como string.
func (d *Duto) GetProxydFim() string {
	return ParaDataString(d.DFim)
}

// SetProxydFim define dFim a partir de uma string.
func (d *Duto) SetProxydFim(valor string) {
	d.DFim, _ = time.Parse("2006-01-02", valor)
}
