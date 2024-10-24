package types

// tpRod representa os tipos de rodados.
type tpRod int

const (
	// NaoAplicavel indica que não se aplica.
	NaoAplicavel tpRod = 0

	// Truck representa um caminhão.
	Truck tpRod = 1

	// Toco representa um caminhão toco.
	Toco tpRod = 2

	// CavaloMecanico representa um cavalo mecânico.
	CavaloMecanico tpRod = 3

	// Van representa uma van.
	Van tpRod = 4

	// Utilitario representa um veículo utilitário.
	Utilitario tpRod = 5

	// Outros representa outros tipos de veículos.
	Outros tpRod = 6
)
