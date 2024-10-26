package delivery

import (
	"ctedos/informacoes/tipos"
	"ctedos/utils"
	"time"
)

// ComData representa uma estrutura de complemento de data.
type ComData struct {
	TpPer tipos.TpPer `json:"tpPer"`
	DProg time.Time   `json:"-"`
}

// ProxydProg retorna a data formatada como string.
func (cd *ComData) ProxydProg() string {
	return utils.ParaDataString(cd.DProg)
}

// SetProxydProg define a data a partir de uma string formatada.
func (cd *ComData) SetProxydProg(value string) error {
	dt, err := time.Parse("2006-01-02", value) // Exemplo de formato de data ISO 8601 (AAAA-MM-DD)
	if err != nil {
		return err
	}
	cd.DProg = dt
	return nil
}
