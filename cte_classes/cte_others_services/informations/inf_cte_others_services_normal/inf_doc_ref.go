package inf_cte_others_services_normal

import (
	"time"
	"encoding/xml"
	"fmt"
)

type InfDocRef struct {
	NDoc      string    `xml:"nDoc"`      // Número do documento
	Serie     *int16    `xml:"serie"`     // Série do documento
	SubSerie  *int16    `xml:"subserie"`  // Sub-série do documento
	DEmi      time.Time `xml:"-"`         // Data de emissão (ignorada no XML direto)
	ProxydEmi string    `xml:"dEmi"`      // Proxy para serializar a data de emissão
	VDoc      *float64  `xml:"vDoc"`      // Valor do documento
}

// Método para serializar a data de emissão no formato esperado
func (i *InfDocRef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias InfDocRef
	if i != nil {
		i.ProxydEmi = i.DEmi.Format("2006-01-02") // Formato AAAA-MM-DD
	}
	return e.EncodeElement((*Alias)(i), start)
}

// Método para deserializar a data de emissão
func (i *InfDocRef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias InfDocRef
	aux := &struct {
		ProxydEmi string `xml:"dEmi"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	parsedDate, err := time.Parse("2006-01-02", aux.ProxydEmi)
	if err != nil {
		return fmt.Errorf("data inválida: %v", err)
	}

	i.DEmi = parsedDate
	return nil
}
