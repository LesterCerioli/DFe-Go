package inf_modals

// DetCont representa detalhes do contêiner.
type DetCont struct {
	NCont  string        `xml:"nCont"`
	Lacre  []Lacre       `xml:"lacre"`
	InfDoc []InfDocAquav `xml:"infDoc"`
}

// Lacre representa o lacre de um contêiner.
type Lacre struct {
	NLacre string `xml:"nLacre"`
}

// InfDocAquav representa as informações do documento aquaviário.
type InfDocAquav struct {
	InfNF  []InfNFAquav  `xml:"infNF"`
	InfNFe []InfNFeAquav `xml:"infNFe"`
}

// InfNFAquav representa as informações da nota fiscal aquaviária.
type InfNFAquav struct {
	Serie   int16    `xml:"serie"`
	NDoc    string   `xml:"nDoc"`
	UnidRat *float64 `xml:"unidRat"`
}

// InfNFeAquav representa as informações da nota fiscal eletrônica aquaviária.
type InfNFeAquav struct {
	Chave   string   `xml:"chave"`
	UnidRat *float64 `xml:"unidRat"`
}

// Funções auxiliares para arredondar valores (adapte se necessário)
func Arredondar(valor *float64, casasDecimais int) *float64 {
	if valor == nil {
		return nil
	}
	fator := math.Pow(10, float64(casasDecimais))
	novoValor := math.Round(*valor*fator) / fator
	return &novoValor
}
