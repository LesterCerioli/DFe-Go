package inf_cte_normal

import (
	"encoding/xml"
)

// Cobr representa as cobranças no CTe.
type Cobr struct {
	Fat Fat   `xml:"fat" json:"fat" order:"1"`
	Dup []Dup `xml:"dup" json:"dup" order:"2"`
}

// Fat representa a estrutura da fatura.
type Fat struct {
	// Defina os campos da fatura aqui
}

// Dup representa a estrutura da duplicata.
type Dup struct {
	// Defina os campos da duplicata aqui
}
