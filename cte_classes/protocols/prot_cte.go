package protocols

import "CTe.Classes.Servicos.Tipos"

// ProtCTe representa o protocolo de resposta do CT-e.
type ProtCTe struct {
	// PR02 - Versão do leiaute das informações de Protocolo.
	Versao TipoVersao `xml:"versao,attr" json:"versao"` // Assumindo que 'versao' é um tipo definido em 'Tipos'

	// PR03 - Informações do Protocolo de resposta. TAG a ser assinada
	InfProt InfProt `xml:"infProt" json:"infProt"`
}

// NewProtCTe cria uma nova instância de ProtCTe.
func NewProtCTe() *ProtCTe {
	return &ProtCTe{
		InfProt: InfProt{},
	}
}
