package schemas

import (
	"encoding/xml"
	"time"
)

// RetInfEvento representa as informações do evento no retorno da solicitação.
type RetInfEvento struct {
	TpAmb        byte      `xml:"tpAmb" json:"tpAmb"`
	VerAplic     string    `xml:"verAplic" json:"verAplic"`
	COrgao       byte      `xml:"cOrgao" json:"cOrgao"`
	CStat        byte      `xml:"cStat" json:"cStat"`
	XMotivo      string    `xml:"xMotivo" json:"xMotivo"`
	ChCTe        string    `xml:"chCTe" json:"chCTe"`
	TpEvento     uint      `xml:"tpEvento" json:"tpEvento"`
	XEvento      string    `xml:"xEvento" json:"xEvento"`
	NSeqEvento   byte      `xml:"nSeqEvento" json:"nSeqEvento"`
	CNPJDest     string    `xml:"CNPJDest" json:"CNPJDest"`
	EmailDest    string    `xml:"emailDest" json:"emailDest"`
	DhRegEvento  time.Time `xml:"-" json:"-"`
	ProxyDhRegEvento string `xml:"dhRegEvento" json:"dhRegEvento"`
	NProt        string    `xml:"nProt" json:"nProt"`
}

// UnmarshalXML é utilizado para fazer o parse do campo dhRegEvento
func (r *RetInfEvento) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias RetInfEvento
	aux := &struct {
		*Alias
		ProxyDhRegEvento string `xml:"dhRegEvento"`
	}{
		Alias: (*Alias)(r),
	}

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	if aux.ProxyDhRegEvento != "" {
		dhRegEvento, err := time.Parse(time.RFC3339, aux.ProxyDhRegEvento)
		if err != nil {
			return err
		}
		r.DhRegEvento = dhRegEvento
	}

	return nil
}
