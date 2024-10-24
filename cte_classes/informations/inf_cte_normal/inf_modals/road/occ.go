package road

import (
	"CTe.Classes.Extensoes"
	"time"
)

// Occ representa informações de ocorrência.
type Occ struct {
	Serie   string   `json:"serie" xml:"serie"`
	NOcc    int      `json:"nOcc" xml:"nOcc"`
	dEmi    time.Time `json:"-" xml:"-"`
	EmiOcc  EmiOcc   `json:"emiOcc" xml:"emiOcc"`
}

// ProxyDEmi retorna a data de emissão no formato de string.
func (o *Occ) ProxyDEmi() string {
	return o.dEmi.ParaDataString()
}

// SetProxyDEmi converte a string para o tipo time.Time.
func (o *Occ) SetProxyDEmi(value string) {
	o.dEmi, _ = time.Parse("2006-01-02", value) // Altere o formato conforme necessário
}
