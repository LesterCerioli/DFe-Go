package types

// respSeg representa os responsáveis pela segurança.
type respSeg string

const (
	// Remetente é o responsável pelo envio.
	Remetente respSeg = "0" // 0 - Remetente
	// Expedidor é o responsável pelo expedição.
	Expedidor respSeg = "1" // 1 - Expedidor
	// Recebedor é o responsável pelo recebimento.
	Recebedor respSeg = "2" // 2 - Recebedor
	// Destinatario é o responsável pela entrega.
	Destinatario respSeg = "3" // 3 - Destinatário
	// EmitenteDoCTe é o emitente do CTe.
	EmitenteDoCTe respSeg = "4" // 4 - Emitente do CTe
	// TomadorDoServico é o tomador do serviço.
	TomadorDoServico respSeg = "5" // 5 - Tomador do Serviço
)
