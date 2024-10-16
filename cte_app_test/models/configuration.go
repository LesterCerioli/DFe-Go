package models

import ()

type Configuration struct {
	Company                     Company                     `json:"company"`
	ConfigureDigitalCertificate ConfigureDigitalCertificate `json:"configure_digital_certificate"`
	ConfigWebService            ConfigWebService            `json:"config_web_service"`
	DiretorioSalvarXml          string                      `json:"diretorio_salvar_xml"`
	IsSalvarXml                 bool                        `json:"is_salvar_xml"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Company:                     Company{},
		ConfigureDigitalCertificate: ConfigureDigitalCertificate{},
		ConfigWebService:            ConfigWebService{},
	}
}
