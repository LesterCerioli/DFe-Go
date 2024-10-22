package delivery

import (
	"ctedos/informacoes/tipos"
	"ctedos/utils"
	"time"
)

// ComHora representa uma estrutura de complemento de hora.
type ComHora struct {
	TpHor tipos.TpHor   `json:"tpHor"`
	HProg time.Duration `json:"-"`
}

// ProxyhProg retorna a hora formatada como string.
func (ch *ComHora) ProxyhProg() string {
	return utils.ParaHoraString(ch.HProg)
}

// SetProxyhProg define a hora a partir de uma string formatada.
func (ch *ComHora) SetProxyhProg(value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	ch.HProg = duration
	return nil
}
