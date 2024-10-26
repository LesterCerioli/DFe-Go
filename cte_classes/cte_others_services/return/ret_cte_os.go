package return

// RetornoBase deve ser definida conforme sua implementação em Go.
type RetornoBase struct {
	// Adicione os campos relevantes aqui
}

// RetCTeOS representa a estrutura equivalente em Go da classe C# retCTeOS.
type RetCTeOS struct {
	Versao    string           `json:"versao"`    // AR02 - Versão do leiaute
	TpAmb     TipoAmbiente     `json:"tpAmb"`     // AR03 - Identificação do Ambiente
	CUF       Estado           `json:"cUF"`       // AR06a - Código da UF que atendeu a solicitação
	VerAplic  string           `json:"verAplic"`  // AR04 - Versão do Aplicativo que recebeu o Lote
	CStat     int              `json:"cStat"`     // AR05 - Código do status da resposta
	AMotivo   string           `json:"xMotivo"`   // AR06 - Descrição literal do status da resposta
	ProtCTe   ProtCTe          `json:"protCTe"`   // AR07 - Dados do Recibo do Lote
	RetornoXmlString string     `json:"retornoXmlString"` // Retorno do XML em string
	EnvioXmlString   string     `json:"envioXmlString"`   // Envio do XML em string
}

// FuncoesXml deve ser definida conforme sua implementação em Go.
var FuncoesXml = struct {
	XmlStringParaClasse func(string) *RetCTeOS
	ClasseParaXmlString func(CTeOSClassesCTeOS) string
}{}

// TipoAmbiente deve ser definida conforme sua implementação em Go.
type TipoAmbiente string

// Estado deve ser definida conforme sua implementação em Go.
type Estado string

// ProtCTe deve ser definida conforme sua implementação em Go.
type ProtCTe struct {
	// Adicione os campos relevantes aqui
}

// LoadXml carrega um XML e retorna uma instância de RetCTeOS.
func LoadXml(xml string) *RetCTeOS {
	retorno := FuncoesXml.XmlStringParaClasse(xml)
	retorno.RetornoXmlString = xml
	return retorno
}

// LoadXml com cteOS carrega um XML e um objeto CTeOSClasses e retorna uma instância de RetCTeOS.
func LoadXmlWithCTeOS(xml string, cteOS CTeOSClassesCTeOS) *RetCTeOS {
	retorno := LoadXml(xml)
	retorno.EnvioXmlString = FuncoesXml.ClasseParaXmlString(cteOS)
	return retorno
}

// LoadXml com cteOS carrega um objeto CTeOSClasses no RetCTeOS.
func (r *RetCTeOS) LoadXmlWithCTeOS(cteOS CTeOSClassesCTeOS) {
	r.EnvioXmlString = FuncoesXml.ClasseParaXmlString(cteOS)
}
