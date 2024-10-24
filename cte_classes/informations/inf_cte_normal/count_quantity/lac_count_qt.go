package count_quantity

// LacContQt representa informações sobre lacres de contêiner.
type LacContQt struct {
	nLacre string // Número do lacre
}

// GetNLacre retorna o número do lacre.
func (l *LacContQt) GetNLacre() string {
	return l.nLacre
}

// SetNLacre define o número do lacre.
func (l *LacContQt) SetNLacre(value string) {
	l.nLacre = value
}
