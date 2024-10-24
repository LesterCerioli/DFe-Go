package inf_modals

import (
	"infCTeNormal/infModals/aereos"
	"time"
)

// Aereo representa o modal aéreo e implementa a interface ContainerModal.
type Aereo struct {
	NMinu           string             `xml:"nMinu"`
	NOCA            string             `xml:"nOCA"`
	DPrevAereo      time.Time          `xml:"-"`
	ProxyDPrevAereo string             `xml:"dPrevAereo"`
	XLAgEmi         string             `xml:"xLAgEmi"`
	IdT             string             `xml:"IdT"`
	NatCarga        aereos.NatCarga    `xml:"natCarga"`
	Tarifa          aereos.Tarifa      `xml:"tarifa"`
	Peri            []aereos.AereoPeri `xml:"peri"`
}

// GetProxyDPrevAereo é um método que retorna a data formatada como string.
func (a *Aereo) GetProxyDPrevAereo() string {
	return a.DPrevAereo.Format("2006-01-02")
}

// SetProxyDPrevAereo é um método que define a data a partir de uma string.
func (a *Aereo) SetProxyDPrevAereo(value string) error {
	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}
	a.DPrevAereo = date
	return nil
}
