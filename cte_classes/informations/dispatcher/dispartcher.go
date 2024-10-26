package dispatcher

// exped representa as informações do expedidor.
type exped struct {
	CNPJ       string     `xml:"CNPJ"`
	CPF        string     `xml:"CPF"`
	IE         string     `xml:"IE"`
	XNome      string     `xml:"xNome"`
	Fone       string     `xml:"fone"`
	EnderExped enderExped `xml:"enderExped"`
	Email      string     `xml:"email"`
}
