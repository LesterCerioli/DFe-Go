package distribution_dfe

import (
	"encoding/xml"
	"time"
)

// RetDistDFeInt representa a resposta da solicitação de distribuição de DF-e.
type RetDistDFeInt struct {
	XMLName xml.Name `xml:"retDistDFeInt"`

	// B02 - Versão do leiaute
	Versao     float64                `xml:"versao,attr"` // Atributo XML
	TpAmb      byte                   `xml:"tpAmb"`       // B03 - Identificação do Ambiente
	VerAplic   string                 `xml:"verAplic"`    // B04 - Versão do aplicativo
	CStat      int                    `xml:"cStat"`       // B05 - Código do status
	EMotivo    string                 `xml:"xMotivo"`     // B06 - Descrição do status
	DhResp     time.Time              `xml:"dhResp"`      // B07 - Data e hora da resposta
	UltNSU     int64                  `xml:"ultNSU"`      // B08 - Último NSU pesquisado
	MaxNSU     int64                  `xml:"maxNSU"`      // B09 - Maior NSU existente
	LoteDistDFeInt []LoteDistDFeInt   `xml:"loteDistDFeInt"` // Conjunto de informações resumidas
}

// LoteDistDFeInt representa o conjunto de informações resumidas e documentos fiscais eletrônicos.
type LoteDistDFeInt struct {
	// Defina os campos conforme necessário
	// Exemplo de campo que poderia estar presente
	NSU    int64  `xml:"NSU,attr"` // Exemplo de atributo
	Schema string `xml:"schema,attr"` // Exemplo de atributo
	XmlNfe []byte `xml:",chardata"` // Conteúdo em base64
}
