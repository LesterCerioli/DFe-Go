package dao

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"cte_app_test_netcore/models"
)

type ConfigDao struct {
    caminhoAplicacao string
    nomeArquivoXml   string
}

func NewConfiguracaoDao() *ConfigDao {
    execPath, err := os.Executable()
    if err != nil {
        return nil
    }
    return &ConfigDao{
        caminhoAplicacao: filepath.Dir(execPath),
        nomeArquivoXml:   "Configuracao.xml",
    }
}

func (dao *ConfigDao) SalvarConfiguracao(configuracao *models.Config) error {
    xmlData, err := xml.MarshalIndent(configuracao, "", "  ")
    if err != nil {
        return err
    }

    filePath := filepath.Join(dao.caminhoAplicacao, dao.nomeArquivoXml)
    return os.WriteFile(filePath, xmlData, 0644)
}

func (dao *ConfigDao) BuscarConfiguracao() (*models.Config, error) {
    filePath := filepath.Join(dao.caminhoAplicacao, dao.nomeArquivoXml)

    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return nil, nil
    }

    xmlData, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var configuracao models.Config
    err = xml.Unmarshal(xmlData, &configuracao)
    if err != nil {
        return nil, err
    }

    return &configuracao, nil
}
