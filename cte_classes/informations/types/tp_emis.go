package types

// tpEmis representa a forma de emissão da CT-e.
type tpEmis int

const (
	// teNormal representa a emissão normal (não em contingência).
	teNormal tpEmis = 1 // 1 - Emissão normal (não em contingência)

	// teEPEC representa a contingência EPEC pela SVC.
	teEPEC tpEmis = 4 // 4 - Contingência EPEC pela SVC

	// teFSDA representa a contingência FS-DA, com impressão do DANFE em formulário de segurança.
	teFSDA tpEmis = 5 // 5 - Contingência FS-DA, com impressão do DANFE em formulário de segurança

	// teSVCRS representa a contingência SVC-RS (SEFAZ Virtual de Contingência do RS).
	teSVCRS tpEmis = 7 // 7 - Contingência SVC-RS (SEFAZ Virtual de Contingência do RS)

	// teSVCSP representa a contingência SVC-SP (SEFAZ Virtual de Contingência de SP).
	teSVCSP tpEmis = 8 // 8 - Contingência SVC-SP (SEFAZ Virtual de Contingência de SP)
)
