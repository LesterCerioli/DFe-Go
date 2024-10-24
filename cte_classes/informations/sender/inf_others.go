package sender

// InfOutros representa informações sobre outros documentos fiscais.
type InfOutros struct {
	TpDoc     int     // Tipo de documento
	DescOutros string  // Descrição de outros documentos
	NDoc      string  // Número do documento
	DEmi      string  // Data de emissão
	VDocFisc  float64 // Valor do documento fiscal
}
