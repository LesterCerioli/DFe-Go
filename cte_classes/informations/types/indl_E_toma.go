package types

// indIEToma representa os indicadores de inscrição no ICMS.
type indIEToma int

const (
	// ContribuinteIcms indica que é um contribuinte de ICMS.
	ContribuinteIcms indIEToma = 1 // "1"

	// ContribuinteIsentoDeInscricao indica que é um contribuinte isento de inscrição.
	ContribuinteIsentoDeInscricao indIEToma = 2 // "2"

	// NaoContribuinte indica que não é contribuinte.
	NaoContribuinte indIEToma = 9 // "9"
)
