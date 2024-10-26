package types

// tpDoc representa os tipos de documentos fiscais.
type tpDoc string

const (
	// Declaracao representa o tipo de documento de declaração.
	Declaracao tpDoc = "00" // 00 - Declaração
	// Dutoviario representa o tipo de documento dutoviário.
	Dutoviario tpDoc = "10" // 10 - Dutoviário
	// CFeSAT representa o tipo de documento CFeSAT.
	CFeSAT tpDoc = "59" // 59 - CFeSAT
	// NFCe representa o tipo de documento NFCe.
	NFCe tpDoc = "65" // 65 - NFCe
	// Outros representa outros tipos de documentos.
	Outros tpDoc = "99" // 99 - Outros
)
