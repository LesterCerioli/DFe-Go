package previous_docs

import (
	"encoding/xml"
	"github.com/yourusername/yourproject/DFe/Classes/Entidades" // Altere para o caminho correto do pacote
)

// EmiDocAnt representa informações sobre documentos anteriores emitidos.
type EmiDocAnt struct {
	CNPJ    string       `xml:"CNPJ"`
	CPF     string       `xml:"CPF"`
	IE      string       `xml:"IE"`
	UF      Estado       `xml:"-"`
	UFProxy string       `xml:"UF"`
	XNome   string       `xml:"xNome"`
	IDDocAnt []IDDocAnt  `xml:"idDocAnt"`
}

// SetUFProxy define o estado a partir da sigla.
func (e *EmiDocAnt) SetUFProxy(value string) {
	e.UF = UF.SiglaParaEstado(value)
}

// GetUFProxy retorna a sigla do estado.
func (e *EmiDocAnt) GetUFProxy() string {
	return e.UF.GetSiglaUfString()
}
