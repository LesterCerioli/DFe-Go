package inf_cte_others_services_normal

import (
	"encoding/xml"
)

// VersaoModal representa a versão do modal
type VersaoModal string

// InfModalOs é a estrutura equivalente em Go da classe C# infModalOs
type InfModalOs struct {
	XMLName      xml.Name       `xml:"infModalOs"`
	VersaoModal  VersaoModal    `xml:"versaoModal,attr"`
	ContainerModal ContainerModal `xml:"rodoOS"`
}

// ContainerModal é a representação genérica do container modal (no caso rodoOS)
type ContainerModal struct {
	XMLName xml.Name `xml:"rodoOS"`
	// Adicione os campos específicos de rodoOS aqui
}
