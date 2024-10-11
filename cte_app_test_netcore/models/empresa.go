package models

// Estado representa a sigla do estado.
type Estado string

// Empresa representa a entidade da empresa para emissão de NFe e NFC-e.
type Empresa struct {
	Cnpj                 string  `json:"cnpj"`
	InscricaoEstadual    string  `json:"inscricaoEstadual"`
	Nome                 string  `json:"nome"`
	NomeFantasia         string  `json:"nomeFantasia"`
	Logradouro           string  `json:"logradouro"`
	Numero               string  `json:"numero"`
	Complemento          string  `json:"complemento"`
	Bairro               string  `json:"bairro"`
	CodigoIbgeMunicipio  int64   `json:"codigoIbgeMunicipio"`
	NomeMunicipio        string  `json:"nomeMunicipio"`
	Cep                  string  `json:"cep"`
	SiglaUf              Estado  `json:"siglaUf"`
	Telefone             string  `json:"telefone"`
	Email                string  `json:"email"`
	RNTRC                string  `json:"rntrc"`
}
