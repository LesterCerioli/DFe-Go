package inf_cte_normal

import (
	"encoding/xml"
)

// InfCarga representa as informações de carga do CTe.
type InfCarga struct {
	VCarga      *float64 `xml:"vCarga"`      // Valor da carga, arredondado para 2 casas decimais
	ProPred     string   `xml:"proPred"`     // Produto da carga
	XOutCat     string   `xml:"xOutCat"`     // Outra categoria
	InfQ        []InfQ   `xml:"infQ"`        // Informações de quantidade
	VCargaAverb *float64 `xml:"vCargaAverb"` // Valor da carga averbada (opcional)
	InfDoc      InfDoc   `xml:"infDoc"`      // Informações do documento
	DocAnt      DocAnt   `xml:"docAnt"`      // Documentos anteriores
}

// VCargaArredondado retorna o valor da carga arredondado para 2 casas decimais.
func (ic *InfCarga) VCargaArredondado() *float64 {
	if ic.VCarga == nil {
		return nil
	}
	// Aqui você pode implementar a lógica de arredondamento desejada.
	roundedValue := round(*ic.VCarga, 2)
	return &roundedValue
}

// VCargaAverbArredondado retorna o valor da carga averbada arredondado para 2 casas decimais.
func (ic *InfCarga) VCargaAverbArredondado() *float64 {
	if ic.VCargaAverb == nil {
		return nil
	}
	// Aqui você pode implementar a lógica de arredondamento desejada.
	roundedValue := round(*ic.VCargaAverb, 2)
	return &roundedValue
}

// Funcão para arredondar valores
func round(value float64, places int) float64 {
	// Lógica de arredondamento
	// Implemente de acordo com suas necessidades
	return value // Substitua pelo cálculo real
}

// VCargaSpecified indica se vCarga está definido.
func (ic *InfCarga) VCargaSpecified() bool {
	return ic.VCarga != nil
}

// VCargaAverbSpecified indica se vCargaAverb está definido.
func (ic *InfCarga) VCargaAverbSpecified() bool {
	return ic.VCargaAverb != nil
}
