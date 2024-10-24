package inf_cte_normal

import (
	"CTe.Classes.Informacoes.infCTeNormal.docAnteriores"
)

// DocAnt representa os documentos anteriores no CTe.
type DocAnt struct {
	EmiDocAnt []docAnteriores.EmiDocAnt `xml:"emiDocAnt"` // Lista de documentos emitidos anteriormente
}
