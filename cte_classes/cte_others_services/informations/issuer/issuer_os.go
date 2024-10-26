package issuer

// EmitOs representa a estrutura equivalente em Go da classe C# emitOs.
type EmitOs struct {
	CNPJ     string   `json:"CNPJ"`     // Cadastro Nacional da Pessoa Jurídica
	IE       string   `json:"IE"`       // Inscrição Estadual
	IEST     string   `json:"IEST,omitempty"` // Inscrição Estadual do Substituto, não é obrigatório
	XNome    string   `json:"xNome"`    // Nome do Emitente
 XFant    string   `json:"xFant"`    // Nome Fantasia
	EnderEmit EnderEmit `json:"enderEmit"` // Endereço do Emitente
}

// EnderEmit deve ser definida conforme a classe C# original.
type EnderEmit struct {
	// Definição dos campos da classe enderEmit (a ser ajustada com base na classe C# original)
}
