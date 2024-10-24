package types

// mod representa os modelos de Nota Fiscal.
type mod string

const (
	// NFModelo01A1eAvulsa representa o modelo 01 da Nota Fiscal.
	NFModelo01A1eAvulsa mod = "01" // "01"
	// NFdeProdutor representa a Nota Fiscal de Produtor.
	NFdeProdutor mod = "04" // "04"
)
