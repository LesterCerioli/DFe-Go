package contr4acts

import "encoding/xml"

type INfeServicoAutorizacao interface {
	INfeServico
	ExecuteZip(nfeDadosMsgZip string) *xml.Node
}
