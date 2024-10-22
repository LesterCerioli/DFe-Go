package complement

// Entrega representa uma estrutura que contém informações de entrega.
type Entrega struct {
	ComData interface{} `json:"comData"` // Usando interface{} para permitir diferentes tipos de comDataBase
	ComHora interface{} `json:"comHora"` // Usando interface{} para permitir diferentes tipos de comHoraBase
}
