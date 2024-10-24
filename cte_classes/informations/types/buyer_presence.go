package types

// PresencaComprador indica a presença do comprador.
type PresencaComprador int

const (
	// PcNao indica que o comprador não está presente.
	PcNao PresencaComprador = 0 // 0 - Não

	// PcPresencial indica que o comprador está presente.
	PcPresencial PresencaComprador = 1 // 1 - Presencial

	// PcInternet indica que a compra foi realizada pela internet.
	PcInternet PresencaComprador = 2 // 2 - Internet

	// PcTeleatendimento indica que a compra foi realizada por teleatendimento.
	PcTeleatendimento PresencaComprador = 3 // 3 - Teleatendimento

	// PcEntregaDomicilio indica que a compra foi feita com entrega em domicílio.
	PcEntregaDomicilio PresencaComprador = 4 // 4 - Entrega em Domicílio

	// PcOutros indica outros tipos de presença.
	PcOutros PresencaComprador = 9 // 9 - Outros
)
