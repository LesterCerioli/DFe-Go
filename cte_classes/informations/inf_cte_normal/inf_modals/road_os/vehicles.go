package road_os

import (
	"encoding/xml"
)

// Veic representa as informações do veículo.
type Veic struct {
	CInt       string     `xml:"cInt"`
	RENAVAM    string     `xml:"RENAVAM"`
	Placa      string     `xml:"placa"`
	Tara       *int       `xml:"tara,omitempty"`
	CapKG      *int       `xml:"capKG,omitempty"`
	CapM3      *int       `xml:"capM3,omitempty"`
	TpProp     *TpProp    `xml:"tpProp,omitempty"`
	TpVeic     *TpVeic    `xml:"tpVeic,omitempty"`
	TpRod      *TpRod     `xml:"tpRod,omitempty"`
	TpCar      *TpCar     `xml:"tpCar,omitempty"`
	UF         Estado     `xml:"UF"`
	Prop       Prop       `xml:"prop"`
}

// GetTaraSpecified retorna true se Tara estiver definido.
func (v *Veic) GetTaraSpecified() bool {
	return v.Tara != nil
}

// GetCapKGSpecified retorna true se CapKG estiver definido.
func (v *Veic) GetCapKGSpecified() bool {
	return v.CapKG != nil
}

// GetCapM3Specified retorna true se CapM3 estiver definido.
func (v *Veic) GetCapM3Specified() bool {
	return v.CapM3 != nil
}

// GetTpPropSpecified retorna true se TpProp estiver definido.
func (v *Veic) GetTpPropSpecified() bool {
	return v.TpProp != nil
}

// GetTpVeicSpecified retorna true se TpVeic estiver definido.
func (v *Veic) GetTpVeicSpecified() bool {
	return v.TpVeic != nil
}

// GetTpRodSpecified retorna true se TpRod estiver definido.
func (v *Veic) GetTpRodSpecified() bool {
	return v.TpRod != nil
}

// GetTpCarSpecified retorna true se TpCar estiver definido.
func (v *Veic) GetTpCarSpecified() bool {
	return v.TpCar != nil
}

// ProxyUF retorna a sigla da UF.
func (v *Veic) ProxyUF() string {
	return v.UF.GetSiglaUfString()
}

// SetProxyUF define a UF com base na sigla.
func (v *Veic) SetProxyUF(sigla string) {
	v.UF = v.UF.SiglaParaEstado(sigla)
}
