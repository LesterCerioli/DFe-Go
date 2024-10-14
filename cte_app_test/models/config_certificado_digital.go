package models

import "encoding/json"

type ConfigCertificadoDigital struct {
	NumeroDeSerie  string `json:"numero_de_serie"`
	CaminhoArquivo string `json:"caminho_arquivo"`
	Senha          string `json:"senha"`
	ManterEmCache  bool   `json:"manter_em_cache"`
}

func main() {
	// Example usage
	config := ConfigCertificadoDigital{
		NumeroDeSerie:  "123456",
		CaminhoArquivo: "/path/to/file",
		Senha:          "password",
		ManterEmCache:  true,
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(config)
	println(string(jsonData))
}
