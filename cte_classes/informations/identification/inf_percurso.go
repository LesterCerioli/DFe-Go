package identification

import (
	"encoding/xml"
	"your_project_path/DFe/Classes/Entidades"
	"your_project_path/DFe/Classes/Extensoes"
)

// infPercurso representa as informações do percurso no CT-e.
type infPercurso struct {
	UFPer      Estado `xml:"-"`
	ProxyUFPer string `xml:"UFPer"`
}

// GetProxyUFPer retorna a sigla da UF.
func (i *infPercurso) GetProxyUFPer() string {
	return i.UFPer.GetSiglaUfString()
}

// SetProxyUFPer define a UF a partir da sigla.
func (i *infPercurso) SetProxyUFPer(value string) {
	i.UFPer = UFPer.SiglaParaEstado(value)
}
