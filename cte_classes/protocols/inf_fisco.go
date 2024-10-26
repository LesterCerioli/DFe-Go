package protocols

// InfFisco representa a estrutura das informações do Fisco.
type InfFisco struct {
	CMsg int    `json:"cMsg" xml:"cMsg"`
	XMsg string `json:"xMsg" xml:"xMsg"`
}
