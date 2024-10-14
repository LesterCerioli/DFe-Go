package models

import "encoding/json"

type Configuracao struct {
	Empresa            Empresa                  `json:"empresa"`
	CertificadoDigital ConfigCertificadoDigital `json:"certificado_digital"`
	ConfigWebService   ConfigWebService         `json:"config_web_service"`
	DiretorioSalvarXml string                   `json:"diretorio_salvar_xml"`
	IsSalvarXml        bool                     `json:"is_salvar_xml"`
}

func NewConfiguracao() *Configuracao {
	return &Configuracao{
		Empresa:            Empresa{},
		CertificadoDigital: ConfigCertificadoDigital{},
		ConfigWebService:   ConfigWebService{},
	}
}
