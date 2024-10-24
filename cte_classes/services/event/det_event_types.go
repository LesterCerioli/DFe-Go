package event

// TipoAutor representa os tipos de autores de eventos.
type TipoAutor int

const (
	// TaEmpresaEmitente - 1 - Empresa Emitente
	TaEmpresaEmitente TipoAutor = 1

	// TaEmpresaDestinataria - 2 - Empresa Destinatária
	TaEmpresaDestinataria TipoAutor = 2

	// TaEmpresa - 3 - Empresa
	TaEmpresa TipoAutor = 3

	// TaFisco - 5 - Fisco
	TaFisco TipoAutor = 5

	// TaRFB - 6 - Receita Federal Brasileira
	TaRFB TipoAutor = 6

	// TaOutrosOrgaos - 9 - Outros Órgãos
	TaOutrosOrgaos TipoAutor = 9
)

// Descricao retorna a descrição associada ao tipo de autor.
func (t TipoAutor) Descricao() string {
	switch t {
	case TaEmpresaEmitente:
		return "Empresa Emitente"
	case TaEmpresaDestinataria:
		return "Empresa Destinatária"
	case TaEmpresa:
		return "Empresa"
	case TaFisco:
		return "Fisco"
	case TaRFB:
		return "RFB"
	case TaOutrosOrgaos:
		return "Outros Órgãos"
	default:
		return "Desconhecido"
	}
}
