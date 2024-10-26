package identification

// tomaCteOs representa os dados do tomador do CT-e OS.
type tomaCteOs struct {
	CNPJ      string    // CNPJ do tomador
	CPF       string    // CPF do tomador
	IE        string    // Inscrição Estadual do tomador
	XNome     string    // Nome do tomador
	XFant     string    // Nome fantasia do tomador
	Fone      string    // Telefone do tomador
	EnderToma enderToma // Endereço do tomador
	Email     string    // Email do tomador
}
