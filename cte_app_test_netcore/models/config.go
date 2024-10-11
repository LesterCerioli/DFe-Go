package models

// Configuracao representa a configuração para emissão de NFe e NFC-e.
type Config struct {
	Empresa              *Empresa                 `json:"empresa"`
	CertificadoDigital    *ConfigCertificadoDigital `json:"certificadoDigital"`
	ConfigWebService      *ConfigWebService         `json:"configWebService"`
	DiretorioSalvarXml    string                   `json:"diretorioSalvarXml"`
	IsSalvarXml           bool                     `json:"isSalvarXml"`
}

// NewConfiguracao cria uma nova instância de Configuracao.
func NewConfig() *Config {
	return &Config{
		Empresa:              NewEmpresa(),
		CertificadoDigital:    NewConfigCertificadoDigital(),
		ConfigWebService:      NewConfigWebService(),
	}
}

// NewEmpresa cria uma nova instância de Empresa.
func NewEmpresa() *Empresa {
	return &Empresa{}
}

// NewConfigCertificadoDigital cria uma nova instância de ConfigCertificadoDigital.
func NewConfigCertificadoDigital() *ConfigCertificadoDigital {
	return &ConfigCertificadoDigital{}
}

// NewConfigWebService cria uma nova instância de ConfigWebService.
func NewConfigWebService() *ConfigWebService {
	return &ConfigWebService{}
}
