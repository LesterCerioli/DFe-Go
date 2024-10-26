package types

// FinalidadeNFe representa a finalidade da emissão da NF-e.
type FinalidadeNFe string

const (
	// FnNormal representa a NFe normal.
	FnNormal FinalidadeNFe = "1"
	// FnComplementar representa a NFe complementar.
	FnComplementar FinalidadeNFe = "2"
	// FnAjuste representa a NFe de ajuste.
	FnAjuste FinalidadeNFe = "3"
	// FnDevolucao representa a devolução/retorno.
	FnDevolucao FinalidadeNFe = "4"
)
