package receiver

// Dest representa as informações do destinatário em um CTe.
type Dest struct {
	CNPJ      string    `xml:"CNPJ"`      // CNPJ do destinatário
	CPF       string    `xml:"CPF"`       // CPF do destinatário
	IE        string    `xml:"IE"`        // Inscrição Estadual
	XNome     string    `xml:"xNome"`     // Nome do destinatário
	Fone      string    `xml:"fone"`      // Telefone do destinatário
	ISUF      string    `xml:"ISUF"`      // Inscrição de Municípios
	EnderDest EnderDest `xml:"enderDest"` // Endereço do destinatário
	Email     string    `xml:"email"`     // Email do destinatário
	LocEnt    LocEnt    `xml:"locEnt"`    // Local de Entrega
}

// EnderDest representa o endereço do destinatário.
type EnderDest struct {
	// Defina os campos apropriados para EnderDest aqui
}

// LocEnt representa o local de entrega.
type LocEnt struct {
	// Defina os campos apropriados para LocEnt aqui
}
