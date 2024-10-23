package inf_cte_normal

import (
	"CTe/Classes/Informacoes/infCTeNormal/infModals" // ajuste o caminho conforme sua estrutura de pacotes
	"encoding/xml"
)

// InfModal representa as informações do modal.
type InfModal struct {
	VersaoModal    versaoModal `xml:"versaoModal,attr"`
	ContainerModal interface{} `xml:",any"` // Usado para suportar diferentes tipos de modal
}

// Definição dos tipos que podem ser usados no ContainerModal
// (rodo, rodoOS, aereo, aquav, ferrov, duto, multimodal)
type Rodo infModals.Rodo
type RodoOS infModals.RodoOS
type Aereo infModals.Aereo
type Aquav infModals.Aquav
type Ferrov infModals.Ferrov
type Duto infModals.Duto
type Multimodal infModals.Multimodal
