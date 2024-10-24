package inf_cte_normal

// Peri representa a estrutura para o período com versão 2.00.
type Peri struct {
	NONU        string `json:"nONU" xml:"nONU"`
	XNomeAE     string `json:"xNomeAE" xml:"xNomeAE"`
	XClaRisco   string `json:"xClaRisco" xml:"xClaRisco"`
	GrEmb       string `json:"grEmb" xml:"grEmb"`
	QTotProd    string `json:"qTotProd" xml:"qTotProd"`
	QVolTipo    string `json:"qVolTipo" xml:"qVolTipo"`
	PontoFulgor string `json:"pontoFulgor" xml:"pontoFulgor"`
}
