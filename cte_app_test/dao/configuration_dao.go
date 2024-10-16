package dao

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfigurationDao struct {
	AplicationWay string
	FileNameXML   string
}

func NewConfigurationDao() *ConfigurationDao {
	AplicationWay, _ := os.Executable()
	return &ConfigurationDao{
		AplicationWay: filepath.Dir(AplicationWay),
		FileNameXML:   "Configuracao.xml",
	}
}

func (dao *ConfigurationDao) SaveConfiguration(configuration *models.Configuration) error {
	filePath := filepath.Join(dao.AplicationWay, dao.FileNameXML)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	err = encoder.Encode(configuration)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ConfigurationDao) BuscarConfiguracao() (*models.Configuration, error) {
	filePath := filepath.Join(dao.AplicationWay, dao.FileNameXML)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var configuration models.Configuration
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(bytes, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
