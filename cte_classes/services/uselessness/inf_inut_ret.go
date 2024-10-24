package uselessness

import (
	"encoding/xml"
	"time"
)

// InfInutRet representa as informações de retorno para a inutilização
type InfInutRet struct {
	Id       string         `xml:"Id,attr"`    // DR04 - Identificador da TAG a ser assinada
	TpAmb    TipoAmbiente   `xml:"tpAmb"`      // DR05 - Identificação do Ambiente
	VerAplic string         `xml:"verAplic"`   // Versão do Aplicativo
	CStat    int            `xml:"cStat"`      // DR07 - Código do status da resposta
	XMotivo  string         `xml:"xMotivo"`    // DR08 - Descrição do status da resposta
	CUF      Estado         `xml:"cUF"`        // DR09 - Código da UF que atendeu a solicitação
	Ano      *int           `xml:"ano"`        // DR10 - Ano de inutilização
	CNPJ     string         `xml:"CNPJ"`       // DR11 - CNPJ do emitente
	Mod      *ModeloDocumento `xml:"mod"`      // DR12 - Modelo da NF-e
	Serie    *int16         `xml:"serie"`      // DR13 - Série da NF-e
	NCTIni   *int64         `xml:"nCTIni"`     // DR14 - Número da NF-e inicial a ser inutilizada
	NCTFin   *int64         `xml:"nCTFin"`     // DR15 - Número da NF-e final a ser inutilizada
	DhRecbto *time.Time     `xml:"dhRecbto"`   // DR16 - Data e hora de processamento
	NProt    string         `xml:"nProt"`      // DR17 - Número do Protocolo de Inutilização
}

// ShouldSerialize retorna true se o campo ano deve ser serializado
func (i *InfInutRet) ShouldSerializeAno() bool {
	return i.Ano != nil
}

// ShouldSerialize retorna true se o campo mod deve ser serializado
func (i *InfInutRet) ShouldSerializeMod() bool {
	return i.Mod != nil
}

// ShouldSerialize retorna true se o campo serie deve ser serializado
func (i *InfInutRet) ShouldSerializeSerie() bool {
	return i.Serie != nil
}

// ShouldSerialize retorna true se o campo nCTIni deve ser serializado
func (i *InfInutRet) ShouldSerializeNCTIni() bool {
	return i.NCTIni != nil
}

// ShouldSerialize retorna true se o campo nCTFin deve ser serializado
func (i *InfInutRet) ShouldSerializeNCTFin() bool {
	return i.NCTFin != nil
}

// ShouldSerialize retorna true se o campo dhRecbto deve ser serializado
func (i *InfInutRet) ShouldSerializeDhRecbto() bool {
	return i.DhRecbto != nil
}
