package sender

// LocColeta representa as informações sobre o local de coleta.
type LocColeta struct {
	CNPJ    string // Cadastro Nacional da Pessoa Jurídica
	CPF     string // Cadastro de Pessoa Física
	XNome   string // Nome
	XLgr    string // Logradouro
	Nro     string // Número
	XCpl    string // Complemento
	XBairro string // Bairro
	CMun    int64  // Código do Município
	XMun    string // Nome do Município
	UF      Estado // Unidade Federativa
}

// Estado é um tipo para representar o estado de forma simplificada.
type Estado string

// Funções para manipulação de Estado (para simular os métodos no C#)

// GetSiglaUfString retorna a sigla do estado como string.
func (e Estado) GetSiglaUfString() string {
	return string(e)
}

// SiglaParaEstado converte uma sigla de estado para o tipo Estado.
func SiglaParaEstado(sigla string) Estado {
	return Estado(sigla)
}
