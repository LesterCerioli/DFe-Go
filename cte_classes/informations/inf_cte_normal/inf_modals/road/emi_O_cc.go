package road

import (
	"CTe.Classes.Extensoes"
	"CTe.Classes.Entidades"
)

// EmiOcc representa as informações do emissor da ocorrência.
type EmiOcc struct {
	CNPJ     string         `json:"CNPJ" xml:"CNPJ"`
	CInt     string         `json:"cInt" xml:"cInt"`
	IE       string         `json:"IE" xml:"IE"`
	UF       Entidades.Estado `json:"-" xml:"-"`
	Fone     string         `json:"fone" xml:"fone"`
}

// ProxyUF retorna a sigla do estado.
func (e *EmiOcc) ProxyUF() string {
	return e.UF.GetSiglaUfString()
}

// SetProxyUF converte a sigla do estado para o tipo Estado.
func (e *EmiOcc) SetProxyUF(value string) {
	e.UF = Entidades.Estado.SiglaParaEstado(value)
}
