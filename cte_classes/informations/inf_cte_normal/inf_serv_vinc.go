package inf_cte_normal

import "encoding/xml"

// InfCTeMultimodal representa a estrutura de infCTeMultimodal.
type InfCTeMultimodal struct {
	ChCTeMultimodal string `xml:"chCTeMultimodal"`
}

// InfServVinc representa os serviços vinculados.
type InfServVinc struct {
	InfCTeMultimodal []InfCTeMultimodal `xml:"infCTeMultimodal"`
}
