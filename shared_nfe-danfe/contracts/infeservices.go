package contracts

import (
	"encoding/xml"
)

type NfeCabecMsg struct {
	// Add necessary fields for NfeCabecMsg
}

type INfeServico interface {
	GetNfeCabecMsg() NfeCabecMsg
	SetNfeCabecMsg(msg NfeCabecMsg)
	Execute(nfeDadosMsg *xml.Node) *xml.Node
}
