package sender

// Rem representa o remetente.
type Rem struct {
	CNPJ      string     // Cadastro Nacional da Pessoa Jurídica
	CPF       string     // Cadastro de Pessoa Física
	IE        string     // Inscrição Estadual
	XNome     string     // Nome
	XFant     string     // Nome Fantasia
	Fone      string     // Telefone
	EnderReme EnderReme  // Endereço do remetente
	Email     string     // E-mail
	LocColeta LocColeta  // Local de coleta (Versão 2.00)
}
