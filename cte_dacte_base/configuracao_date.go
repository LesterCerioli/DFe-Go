package configuracao_date


import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

type ConfiguracaoDacte struct {
	Logomarca              []byte
	DocumentoCancelado     bool
	QuebrarLinhasObservacao bool
	Desenvolvedor          string
}

func NewConfiguracaoDacte() *ConfiguracaoDacte {
	return &ConfiguracaoDacte{
		Logomarca:              nil,
		DocumentoCancelado:     false,
		QuebrarLinhasObservacao: false,
	}
}

func (c *ConfiguracaoDacte) ObterLogo() (image.Image, error) {
	if c.Logomarca == nil {
		return nil, nil
	}
	
	img, _, err := image.Decode(bytes.NewReader(c.Logomarca))
	if err != nil {
		return nil, err
	}
	
	return img, nil
}

