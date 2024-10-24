package schemas

import "time"

// MDFe representa os dados do manifesto de documentos fiscais eletrônicos.
type MDFe struct {
	ChMDFe   string    `json:"chMDFe"`
	Modal    string    `json:"modal"`
	DhEmi    time.Time `json:"dhEmi"`
	NProt    string    `json:"nProt"`
	DhRecbto time.Time `json:"dhRecbto"`
}
