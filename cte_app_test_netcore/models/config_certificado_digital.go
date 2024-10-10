package models

import "encoding/json"

type ConfigCertificadoDigital struct {
	NumeroDeSerie   string `json:"numero_de_serie"`
	CaminhoArquivo  string `json:"caminho_arquivo"`
	Senha           string `json:"senha"`
	ManterEmCache   bool   `json:"manter_em_cache"`
}

func (c *ConfigCertificadoDigital) ToJSON() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

