package types

// tpServ representa os tipos de serviço.
type tpServ int

const (
	// normal representa um serviço normal.
	normal tpServ = 0

	// subcontratacao representa um serviço de subcontratação.
	subcontratacao tpServ = 1

	// redespacho representa um serviço de redespacho.
	redespacho tpServ = 2

	// redespachoIntermediario representa um redespacho intermediário.
	redespachoIntermediario tpServ = 3

	// servicoVinculadoMultimodal representa um serviço vinculado multimodal.
	servicoVinculadoMultimodal tpServ = 4

	// transportePessoas representa um serviço de transporte de pessoas.
	transportePessoas tpServ = 6

	// transporteValores representa um serviço de transporte de valores.
	transporteValores tpServ = 7

	// excessoBagagem representa um serviço de excesso de bagagem.
	excessoBagagem tpServ = 8
)
