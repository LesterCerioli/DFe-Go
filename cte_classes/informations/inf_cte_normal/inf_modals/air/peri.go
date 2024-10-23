package air

// AereoPeri representa as informações sobre a carga aérea.
type AereoPeri struct {
	NONU   string   `json:"nONU" xml:"nONU"`
	QTotEmb string   `json:"qTotEmb" xml:"qTotEmb"`
	InfTotAP infTotAP `json:"infTotAP" xml:"infTotAP"`
}
