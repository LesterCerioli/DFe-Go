package types

// IdentificadorOperacaoCsc representa o tipo de operação do CSC.
type IdentificadorOperacaoCsc int

const (
	// IoConsultaCscAtivos indica consulta de CSC ativos.
 IoConsultaCscAtivos IdentificadorOperacaoCsc = 1 // Consulta CSC Ativos

	// IoSolicitaNovoCsc indica solicitação de novo CSC.
 IoSolicitaNovoCsc IdentificadorOperacaoCsc = 2 // Solicita novo CSC

	// IoRevogaCscAtivo indica revogação de CSC ativo.
 IoRevogaCscAtivo IdentificadorOperacaoCsc = 3 // Revoga CSC Ativo
)
