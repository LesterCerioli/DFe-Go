package types

// tpTraf representa os tipos de tráfego.
type tpTraf int

const (
	// Proprio representa tráfego próprio.
	Proprio tpTraf = 0

	// Mutuo representa tráfego mútuo.
	Mutuo tpTraf = 1

	// Rodoferroviario representa tráfego rodoferroviário.
	Rodoferroviario tpTraf = 2

	// Rodoviario representa tráfego rodoviário.
	Rodoviario tpTraf = 3
)
