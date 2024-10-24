package distribution_dfe

import (
	"encoding/base64"
	"encoding/xml"
)

// LoteDistDFeInt representa o conjunto de informações resumidas e documentos fiscais eletrônicos de interesse da pessoa ou empresa.
type LoteDistDFeInt struct {
	XMLName xml.Name `xml:"loteDistDFeInt"`

	// NSU do documento fiscal
	NSU     int64  `xml:"NSU,attr"` // Atributo XML
	Schema  string `xml:"schema,attr"` // Atributo XML

	// Conteúdo da Tag schema, compactado no padrão gZip.
	XmlNfe []byte `xml:",chardata"` // Conteúdo base64Binary
}

// DecodeXmlNfe decodifica o campo XmlNfe de base64.
func (l *LoteDistDFeInt) DecodeXmlNfe() ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(l.XmlNfe))
}
