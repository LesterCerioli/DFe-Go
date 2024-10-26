package ext

import (
	"ctedos/servicos/tipos"
	"errors"
	"fmt"
)

// GetString retorna a string correspondente à versão do CT-e.
func GetString(versao tipos.Versao) (string, error) {
	switch versao {
	case tipos.Ve200:
		return "2.00", nil
	case tipos.Ve300:
		return "3.00", nil
	case tipos.Ve400:
		return "4.00", nil
	default:
		return "", errors.New("versão de CT-e inválida. Para emissão, apenas as versões 2.00, 3.00 e 4.00 são aceitas")
	}
}
