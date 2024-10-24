package event

import "encoding/xml"

// EvCCeCTe representa o evento de Carta de Correção (evCCeCTe)
type EvCCeCTe struct {
	XMLName     xml.Name       `xml:"evCCeCTe"`
	DescEvento  string         `xml:"descEvento"`
	InfCorrecao []InfCorrecao  `xml:"infCorrecao"`
	XCondUso    string         `xml:"xCondUso"`
}

// InfCorrecao representa a correção feita em um campo específico
type InfCorrecao struct {
	GrupoAlterado    string  `xml:"grupoAlterado"`
	CampoAlterado    string  `xml:"campoAlterado"`
	ValorAlterado    string  `xml:"valorAlterado"`
	NroItemAlterado  *int    `xml:"nroItemAlterado,omitempty"` // Campo opcional
}

// NewEvCCeCTe cria uma nova instância de EvCCeCTe com valores padrão
func NewEvCCeCTe() *EvCCeCTe {
	return &EvCCeCTe{
		DescEvento: "Carta de Correcao",
		XCondUso: "A Carta de Correcao e disciplinada pelo Art. 58-B do CONVENIO/SINIEF 06/89: " +
			"Fica permitida a utilizacao de carta de correcao, para regularizacao de erro ocorrido na emissao de documentos " +
			"fiscais relativos a prestacao de servico de transporte, desde que o erro nao esteja relacionado com: I - as variaveis " +
			"que determinam o valor do imposto tais como: base de calculo, aliquota, diferenca de preco, quantidade, valor da prestacao; " +
			"II - a correcao de dados cadastrais que implique mudanca do emitente, tomador, remetente ou do destinatario; " +
			"III - a data de emissao ou de saida.",
	}
}
