package complement

type ComplOs struct {
	XCaracAd  string     `xml:"xCaracAd"`
	XCaracSer string     `xml:"xCaracSer"`
	XEmi      string     `xml:"xEmi"`
	XObs      string     `xml:"xObs"`
	ObsCont   []ObsCont  `xml:"ObsCont"`
	ObsFisco  []ObsFisco `xml:"ObsFisco"`
}

type ObsCont struct {
	// Defina os campos de ObsCont conforme necessário
}

type ObsFisco struct {
	// Defina os campos de ObsFisco conforme necessário
}
