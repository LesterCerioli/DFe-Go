package types

// tpCar representa os tipos de veículos.
type tpCar string

const (
	// NaoAplicavel indica que o tipo de veículo não é aplicável.
	NaoAplicavel tpCar = "00" // 00 - Não Aplicável
	// Aberta indica um veículo tipo aberto.
	Aberta tpCar = "01" // 01 - Aberta
	// Fechada indica um veículo tipo fechado.
	Fechada tpCar = "02" // 02 - Fechada
	// Graneleira indica um veículo tipo graneleira.
	Graneleira tpCar = "03" // 03 - Graneleira
	// PortaContainer indica um veículo tipo porta-container.
	PortaContainer tpCar = "04" // 04 - Porta Container
	// Sider indica um veículo tipo sider.
	Sider tpCar = "05" // 05 - Sider
)
