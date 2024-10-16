package models

import "encoding/json"

type Company struct {
	Cnpj                string `json:"cnpj"`
	InscricaoEstadual   string `json:"inscricaoEstadual"`
	Nome                string `json:"nome"`
	NomeFantasia        string `json:"nomeFantasia"`
	Logradouro          string `json:"logradouro"`
	Numero              string `json:"numero"`
	Complemento         string `json:"complemento"`
	Bairro              string `json:"bairro"`
	CodigoIbgeMunicipio int64  `json:"codigoIbgeMunicipio"`
	NomeMunicipio       string `json:"nomeMunicipio"`
	Cep                 string `json:"cep"`
	SiglaUf             State  `json:"siglaUf"`
	Telefone            string `json:"telefone"`
	Email               string `json:"email"`
	RNTRC               string `json:"rntrc"`
}

// Assuming Estado is an enum-like type, you might define it like this:
type State string

const (
	AC State = "AC"
	AL State = "AL"
	// ... other states ...
)

// MarshalJSON and UnmarshalJSON methods for custom JSON serialization
func (e Company) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Company
	}{
		Company: e,
	})
}

func (e *Company) UnmarshalJSON(data []byte) error {
	type Alias Company
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	return json.Unmarshal(data, &aux)
}
