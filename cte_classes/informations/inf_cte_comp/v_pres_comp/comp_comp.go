package v_pres_comp

// compComp representa a composição de um componente.
type compComp struct {
	xNome string  `xml:"xNome,omitempty"` // Nome do componente
	vComp  float64 `xml:"vComp,omitempty"`  // Valor do componente
}

// NewCompComp cria uma nova instância de compComp.
func NewCompComp(xNome string, vComp float64) *compComp {
	return &compComp{
		xNome: xNome,
		vComp:  vComp,
	}
}

// GetXNome retorna o nome do componente.
func (c *compComp) GetXNome() string {
	return c.xNome
}

// SetXNome define o nome do componente.
func (c *compComp) SetXNome(xNome string) {
	c.xNome = xNome
}

// GetVComp retorna o valor do componente.
func (c *compComp) GetVComp() float64 {
	return c.vComp
}

// SetVComp define o valor do componente.
func (c *compComp) SetVComp(vComp float64) {
	c.vComp = vComp
}
