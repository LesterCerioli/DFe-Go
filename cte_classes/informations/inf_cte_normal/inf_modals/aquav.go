package inf_modals

import (
	"fmt"
	"infCTeNormal/tipos"
)

// Aquav representa o modal aquaviário e implementa a interface ContainerModal.
type Aquav struct {
	VPrest   float64      `xml:"vPrest"`
	VAFRMM   float64      `xml:"vAFRMM"`
	NBooking string       `xml:"nBooking"`
	NCtrl    string       `xml:"nCtrl"`
	XNavio   string       `xml:"xNavio"`
	Balsa    []Balsa      `xml:"balsa"`
	NViag    string       `xml:"nViag"`
	Direc    string       `xml:"direc"`
	PrtEmb   string       `xml:"prtEmb"`
	PrtTrans string       `xml:"prtTrans"`
	PrtDest  string       `xml:"prtDest"`
	TpNav    *tipos.TpNav `xml:"tpNav"`
	Irin     string       `xml:"irin"`
	DetCont  []DetCont    `xml:"detCont"`
}

// SetVPrest define o valor arredondado de vPrest.
func (a *Aquav) SetVPrest(value float64) {
	a.VPrest = round(value, 2)
}

// SetVAFRMM define o valor arredondado de vAFRMM.
func (a *Aquav) SetVAFRMM(value float64) {
	a.VAFRMM = round(value, 2)
}

// TpNavSpecified retorna verdadeiro se TpNav tiver valor.
func (a *Aquav) TpNavSpecified() bool {
	return a.TpNav != nil
}

// Função de arredondamento para dois decimais.
func round(value float64, precision int) float64 {
	format := fmt.Sprintf("%%.%df", precision)
	strValue := fmt.Sprintf(format, value)
	rounded, _ := strconv.ParseFloat(strValue, 64)
	return rounded
}
