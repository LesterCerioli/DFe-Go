package return

// CteOSProc representa a estrutura equivalente em Go da classe C# cteOSProc.
type CteOSProc struct {
	Versao  VersaoServico                 `json:"versao"`  // Versão do serviço
	CTeOS   CTeOS                         `json:"cteOS"`   // CTeOS
	ProtCTe ProtCTe                       `json:"protCTe"` // Protocolo do CTe
}

// VersaoServico define a versão do serviço (defina conforme necessário).
type VersaoServico string

// CTeOS define a estrutura do CTeOS (defina conforme necessário).
type CTeOS struct {
	// Adicione os campos relevantes aqui
}

// ProtCTe define a estrutura do protocolo do CTe (defina conforme necessário).
type ProtCTe struct {
	// Adicione os campos relevantes aqui
}
