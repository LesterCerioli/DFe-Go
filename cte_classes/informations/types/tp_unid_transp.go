package types

// tpUnidTransp representa os tipos de unidade de transporte.
type tpUnidTransp int

const (
	// RodoviarioTracao representa a unidade de transporte rodoviário de tração.
	RodoviarioTracao tpUnidTransp = 1
	
	// RodoviarioReboque representa a unidade de transporte rodoviário de reboque.
	RodoviarioReboque tpUnidTransp = 2
	
	// Navio representa a unidade de transporte por navio.
	Navio tpUnidTransp = 3
	
	// Balsa representa a unidade de transporte por balsa.
	Balsa tpUnidTransp = 4
	
	// Aeronave representa a unidade de transporte por aeronave.
	Aeronave tpUnidTransp = 5
	
	// Vagao representa a unidade de transporte por vagão.
	Vagao tpUnidTransp = 6
	
	// Outros representa outras unidades de transporte.
	Outros tpUnidTransp = 7
)
