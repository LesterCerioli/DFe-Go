package inf_modals

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// Ferrov representa o modal ferroviário.
type Ferrov struct {
	TpTraf   string     `xml:"tpTraf"`
	TrafMut  TrafMut    `xml:"trafMut"`
	Fluxo    string     `xml:"fluxo"`
	IdTrem   string     `xml:"idTrem"`
	VFrete   float64    `xml:"vFrete"`
	FerroEnv []FerroEnv `xml:"ferroEnv"`
	DetVag   []DetVag   `xml:"detVag"`
}

// TrafMut representa o tráfego mútuo ferroviário.
type TrafMut struct {
	RespFat          string     `xml:"respFat"`
	FerrEmi          string     `xml:"ferrEmi"`
	VFrete           *float64   `xml:"vFrete"`
	ChCTeFerroOrigem string     `xml:"chCTeFerroOrigem"`
	FerroEnv         []FerroEnv `xml:"ferroEnv"`
	Fluxo            string     `xml:"fluxo"`
}

// FerroEnv representa o envio ferroviário.
type FerroEnv struct {
	CNPJ       string     `xml:"CNPJ"`
	CInt       string     `xml:"cInt"`
	IE         string     `xml:"IE"`
	XNome      string     `xml:"xNome"`
	EnderFerro EnderFerro `xml:"enderFerro"`
}

// EnderFerro representa o endereço ferroviário.
type EnderFerro struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	CEP     string `xml:"CEP"`
	UF      string `xml:"UF"`
}

// DetVag representa o detalhe de vagões ferroviários.
type DetVag struct {
	NVag   string   `xml:"nVag"`
	Cap    *float64 `xml:"cap"`
	TpVag  string   `xml:"tpVag"`
	PesoR  float64  `xml:"pesoR"`
	PesoBC float64  `xml:"pesoBC"`
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

// SetCap define o valor arredondado de Cap.
func (d *DetVag) SetCap(valor *float64) {
	d.Cap = Arredondar(valor, 3)
}

// SetVFrete define o valor arredondado de vFrete em Ferrov.
func (f *Ferrov) SetVFrete(valor float64) {
	f.VFrete = *Arredondar(&valor, 2)
}

// SetVFreteTrafMut define o valor arredondado de vFrete em TrafMut.
func (t *TrafMut) SetVFreteTrafMut(valor *float64) {
	t.VFrete = Arredondar(valor, 2)
}

// SetCEP define o CEP com zeros à esquerda.
func (e *EnderFerro) SetCEP(valor string) {
	if len(valor) < 8 {
		e.CEP = fmt.Sprintf("%08s", valor)
	} else {
		e.CEP = valor
	}
}

// GetCEP retorna o valor do CEP como string formatada.
func (e *EnderFerro) GetCEP() string {
	cep, _ := strconv.ParseInt(e.CEP, 10, 64)
	return fmt.Sprintf("%08d", cep)
}
