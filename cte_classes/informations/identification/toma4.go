package identification

import "CTe.Classes.Informacoes.Tipos"

// toma4 representa a estrutura do toma4 em Go.
type toma4 struct {
	toma      tipos.Toma // Supondo que 'toma' é um tipo definido no pacote 'Tipos'
	CNPJ      string
	CPF       string
	IE        string
	xNome     string
	xFant     string
	fone      string
	enderToma enderToma // Certifique-se de que enderToma esteja definido
	email     string
}

// NewToma4 cria uma nova instância de toma4 com valores padrão.
func NewToma4() *toma4 {
	return &toma4{
		toma: tipos.Outros, // Inicializa toma com o valor padrão
	}
}
