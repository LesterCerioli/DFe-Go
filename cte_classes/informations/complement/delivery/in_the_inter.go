package delivery

import (
	"ctedos/informacoes/tipos"
	"ctedos/utils"
	"time"
)

// NoInter representa uma estrutura de intervalo de horas.
type NoInter struct {
	TpHor tipos.TpHor `json:"tpHor"`
	HIni  time.Duration `json:"-"`
	HFim  time.Duration `json:"-"`
}

// NewNoInter cria uma nova instância de NoInter com o valor padrão de tpHor.
func NewNoInter() *NoInter {
	return &NoInter{
		TpHor: tipos.NoIntervaloDeTempo,
	}
}

// ProxyhIni retorna o valor de hIni formatado como string.
func (ni *NoInter) ProxyhIni() string {
	return utils.ParaHoraString(ni.HIni)
}

// SetProxyhIni define o valor de hIni a partir de uma string formatada.
func (ni *NoInter) SetProxyhIni(value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	ni.HIni = duration
	return nil
}

// ProxyhFim retorna o valor de hFim formatado como string.
func (ni *NoInter) ProxyhFim() string {
	return utils.ParaHoraString(ni.HFim)
}

// SetProxyhFim define o valor de hFim a partir de uma string formatada.
func (ni *NoInter) SetProxyhFim(value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	ni.HFim = duration
	return nil
