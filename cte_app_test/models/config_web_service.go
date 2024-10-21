package models

import (
	"_/C_/Users/apoll/OneDrive/Documentos/repositories/github/DFe-Go/cte_app_test/models"
)

type ConfigWebService struct {
	UfEmitente     models.State
	Ambiente       Flags.TipoAmbiente
	Serie          int16
	Numeracao      int64
	Versao         Tipos.Versao
	CaminhoSchemas string
	TimeOut        int
}
