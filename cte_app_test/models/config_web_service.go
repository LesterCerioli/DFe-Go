package models

import (
	"CTe/Classes/Servicos/Tipos"
	"DFe/Classes/Entidades"
	"DFe/Classes/Flags"
)

type ConfigWebService struct {
	UfEmitente     Entidades.Estado
	Ambiente       Flags.TipoAmbiente
	Serie          int16
	Numeracao      int64
	Versao         Tipos.Versao
	CaminhoSchemas string
	TimeOut        int
}
