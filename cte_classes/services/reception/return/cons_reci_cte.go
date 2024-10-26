package return

import (
	"encoding/xml"
	"CTe/Classes/Servicos/Tipos" // Ajuste o caminho de importação conforme necessário
)

// ConsReciCTe representa a estrutura de consulta do recibo do CTe
type ConsReciCTe struct {
	Versao   Versao                 `xml:"versao,attr"` // BP02 - Versão do leiaute
	TpAmb    TipoAmbiente           `xml:"tpAmb"`       // BP03 - Identificação do Ambiente
	NRec     string                 `xml:"nRec"`        // BP04 - Número do Recibo
}

// Versao e TipoAmbiente devem ser definidos conforme necessário
// Exemplos de como poderiam ser definidos:
type Versao string

type TipoAmbiente int

const (
	Produção TipoAmbiente = 1
	Homologação TipoAmbiente = 2
)
