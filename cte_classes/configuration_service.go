package cte_classes

import (
	"CTe.Classes.Informacoes/Tipos"
	"CTe.Classes/Servicos/Tipos"
	"crypto/x509"
	"errors"
	"os"
	"sync"
)

type ConfiguracaoServico struct {
	mu                      sync.Mutex
	diretorioSchemas        string
	ConfiguracaoCertificado ConfiguracaoCertificado
	certificado             *x509.Certificate
	TimeOut                 int
	cUF                     Tipos.Estado
	tpAmb                   Tipos.TipoAmbiente
	VersaoLayout            Tipos.Versao
	IsAdicionaQrCode        bool
	IsSalvarXml             bool
	DiretorioSalvarXml      string
	IsValidaSchemas         bool
	tipoEmissao             Tipos.tpEmis
}

var (
	instancia *ConfiguracaoServico
	once      sync.Once
)

// NewConfiguracaoServico cria uma nova instância de ConfiguracaoServico.
func NewConfiguracaoServico() *ConfiguracaoServico {
	return &ConfiguracaoServico{
		ConfiguracaoCertificado: ConfiguracaoCertificado{},
		tipoEmissao:             Tipos.tpEmisNormal, // Assumindo que esse valor é correto
		IsValidaSchemas:         true,
	}
}

// Instancia retorna a instância singleton de ConfiguracaoServico.
func Instancia() *ConfiguracaoServico {
	once.Do(func() {
		instancia = NewConfiguracaoServico()
	})
	return instancia
}

// ObterCertificado obtém o certificado digital.
func (c *ConfiguracaoServico) ObterCertificado() (*x509.Certificate, error) {
	if c.certificado != nil {
		if !c.ConfiguracaoCertificado.ManterDadosEmCache {
			c.certificado.Reset()
		}
	}
	c.certificado = ObterCertificado(c.ConfiguracaoCertificado)
	return c.certificado, nil
}

// Dispose libera os recursos utilizados.
func (c *ConfiguracaoServico) Dispose() {
	if !c.ConfiguracaoCertificado.ManterDadosEmCache && c.certificado != nil {
		c.certificado.Reset()
		c.certificado = nil
	}
}

// SetDiretorioSchemas valida e configura o diretório dos schemas.
func (c *ConfiguracaoServico) SetDiretorioSchemas(diretorio string) error {
	if diretorio != "" {
		if _, err := os.Stat(diretorio); os.IsNotExist(err) {
			return errors.New("Diretório " + diretorio + " não encontrado!")
		}
	}
	c.diretorioSchemas = diretorio
	return nil
}

// ObterVersaoLayoutValida retorna a versão válida do layout.
func (c *ConfiguracaoServico) ObterVersaoLayoutValida() (Tipos.Versao, error) {
	switch c.VersaoLayout {
	case Tipos.VersaoVe200:
		return Tipos.VersaoVe200, nil
	case Tipos.VersaoVe300:
		return Tipos.VersaoVe300, nil
	case Tipos.VersaoVe400:
		return Tipos.VersaoVe400, nil
	default:
		return "", errors.New("versão do layout inválida")
	}
}

// NaoSalvarXml verifica se o XML não deve ser salvo.
func (c *ConfiguracaoServico) NaoSalvarXml() bool {
	return !c.IsSalvarXml
}
