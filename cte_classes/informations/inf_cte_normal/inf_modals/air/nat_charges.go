package air

// NatCarga representa as informações sobre a natureza da carga.
type NatCarga struct {
	XDime    string   `json:"xDime" xml:"xDime"`
	CInfManu []string `json:"cInfManu" xml:"cInfManu"`
	CImp     []string `json:"cIMP" xml:"cIMP"`
}
