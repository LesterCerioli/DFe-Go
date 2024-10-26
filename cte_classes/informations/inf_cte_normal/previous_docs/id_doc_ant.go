package previous_docs

import "encoding/xml"

// IDDocAnt representa informações sobre documentos anteriores.
type IDDocAnt struct {
	IDDocAntPap []IDDocAntPap `xml:"idDocAntPap"`
	IDDocAntEle []IDDocAntEle `xml:"idDocAntEle"`
}
