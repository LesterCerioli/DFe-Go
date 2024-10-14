package dao

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"cte_app_test/entidades"
)

type ConfiguracaoDao struct {
	caminhoAplicacao string
	nomeArquivoXml   string
}

func NewConfiguracaoDao() *ConfiguracaoDao {
	caminhoAplicacao, _ := os.Executable()
	return &ConfiguracaoDao{
		caminhoAplicacao: filepath.Dir(caminhoAplicacao),
		nomeArquivoXml:   "Configuracao.xml",
	}
}

func (dao *ConfiguracaoDao) SalvarConfiguracao(configuracao *entidades.Configuracao) error {
	filePath := filepath.Join(dao.caminhoAplicacao, dao.nomeArquivoXml)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	err = encoder.Encode(configuracao)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ConfiguracaoDao) BuscarConfiguracao() (*entidades.Configuracao, error) {
	filePath := filepath.Join(dao.caminhoAplicacao, dao.nomeArquivoXml)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var configuracao entidades.Configuracao
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(bytes, &configuracao)
	if err != nil {
		return nil, err
	}

	return &configuracao, nil
}
