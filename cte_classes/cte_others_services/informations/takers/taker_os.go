package takers

// TomaOs representa a estrutura equivalente em Go da classe C# tomaOs.
type TomaOs struct {
	CNPJ      string      `json:"CNPJ"`      // CNPJ
	CPF       string      `json:"CPF"`       // CPF
	IE        string      `json:"IE"`        // Inscrição Estadual
	XNome     string      `json:"xNome"`     // Nome
	XFant     string      `json:"xFant"`     // Nome Fantasia
	Fone      string      `json:"fone"`      // Telefone
	EnderToma *EnderTomaOs `json:"enderToma"` // Endereço do Tomador
	Email     string      `json:"email"`     // E-mail
}
