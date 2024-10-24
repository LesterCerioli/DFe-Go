package sender

// LocRet representa o local de retirada.
type LocRet struct {
	CNPJ    string // Cadastro Nacional da Pessoa Jurídica
	CPF     string // Cadastro de Pessoa Física
	XNome   string // Nome
	XLgr    string // Logradouro
	Nro     string // Número
	XCpl    string // Complemento
	XBairro string // Bairro
	CMun    int    // Código do Município
	XMun    string // Nome do Município
	UF      string // Unidade Federativa
}
