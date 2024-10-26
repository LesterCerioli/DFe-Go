package inf_cte_normal

import (
	// Importa o pacote correto para lacContQt, que você deve definir
	"CTe.Classes.Informacoes.infCTeNormal.contQantidades"
)

// ContQt representa as informações de contagem no CTe.
type ContQt struct {
	LacContQt contQantidades.LacContQt // Importa o tipo lacContQt
	NCont     string                   // Número do contrato
	DPrev     string                   // Data prevista
}
