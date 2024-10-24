package inf_cte_normal

import "encoding/xml"

// InfQOs representa a estrutura para infQOs.
type InfQOs struct {
	// Defina os campos de acordo com a estrutura de infQOs.
}

// InfServico representa a estrutura de serviços.
type InfServico struct {
	XDescServ string `xml:"xDescServ"`
	InfQ      InfQOs `xml:"infQ"`
}
