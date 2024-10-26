package types

// toma representa as opções de responsabilidade na operação.
type toma string

const (
	// Remetente indica que a responsabilidade é do remetente.
	Remetente toma = "0" // 0 - Remetente
	// Expedidor indica que a responsabilidade é do expedidor.
	Expedidor toma = "1" // 1 - Expedidor
	// Recebedor indica que a responsabilidade é do recebedor.
	Recebedor toma = "2" // 2 - Recebedor
	// Destinatario indica que a responsabilidade é do destinatário.
	Destinatario toma = "3" // 3 - Destinatário
	// Outros indica que a responsabilidade é de outras partes.
	Outros toma = "4" // 4 - Outros
)
