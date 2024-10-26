package inf_modals

// RodoOS representa as informações do transporte rodoviário OS.
type RodoOS struct {
	TAF            string        `xml:"TAF"`
	NroRegEstadual string        `xml:"NroRegEstadual"`
	Veic           Veic          `xml:"veic"`
	InfFretamento  InfFretamento `xml:"infFretamento"`
}

// Defina a struct Veic com suas propriedades conforme necessário.
type Veic struct {
	// Definições do tipo Veic
}

// Defina a struct InfFretamento com suas propriedades conforme necessário.
type InfFretamento struct {
	// Definições do tipo InfFretamento
}
