package delivery

import (
	"ctedos/informacoes/tipos"
	"ctedos/utils"
	"time"
)

// NoPeriodo representa uma estrutura de período de datas.
type NoPeriodo struct {
	TpPer tipos.TpPer `json:"tpPer"`
	DIni  time.Time   `json:"-"`
	DFim  time.Time   `json:"-"`
}

// NewNoPeriodo cria uma nova instância de NoPeriodo com o valor padrão de tpPer.
func NewNoPeriodo() *NoPeriodo {
	return &NoPeriodo{
		TpPer: tipos.NoPeriodo,
	}
}

// ProxyDIni retorna o valor de DIni formatado como string.
func (np *NoPeriodo) ProxyDIni() string {
	return utils.ParaDataString(np.DIni)
}

// SetProxyDIni define o valor de DIni a partir de uma string formatada.
func (np *NoPeriodo) SetProxyDIni(value string) error {
	date, err := time.Parse("2006-01-02", value) // Utiliza o formato de data padrão "YYYY-MM-DD"
	if err != nil {
		return err
	}
	np.DIni = date
	return nil
}

// ProxyDFim retorna o valor de DFim formatado como string.
func (np *NoPeriodo) ProxyDFim() string {
	return utils.ParaDataString(np.DFim)
}

// SetProxyDFim define o valor de DFim a partir de uma string formatada.
func (np *NoPeriodo) SetProxyDFim(value string) error {
	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}
	np.DFim = date
	return nil
}
