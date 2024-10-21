package extensions

import (
	"os"
	"path/filepath"
	"strings"
	"time"
	"cteservicos/constants"
	"cteservicos/utils"
)

// CarregarDeXmlString converte uma string XML para um objeto retCTeOS
func CarregarDeXmlString(xmlString string) (*retCTeOS, error) {
	return utils.XmlStringParaClasse[retCTeOS](xmlString)
}

// ObterXmlString converte o objeto retCTeOS para uma string no formato XML
func ObterXmlString(retEnviCte *retCTeOS) (string, error) {
	return utils.ClasseParaXmlString(retEnviCte)
}

// SalvarXmlEmDisco salva o objeto retCTeOS como um arquivo XML no disco
func SalvarXmlEmDisco(retEnviCte *retCTeOS, config *configuration_service) error {
	if config.NaoSalvarXml() {
		return nil
	}

	dataEnvio := time.Now()

	if retEnviCte != nil && retEnviCte.protCTe != nil && retEnviCte.protCTe.infProt != nil && retEnviCte.protCTe.infProt.dhRecbto != nil {
		dataEnvio = *retEnviCte.protCTe.infProt.dhRecbto
	}

	caminhoXml := config.DiretorioSalvarXml

	protocolo := "000000"
	if retEnviCte != nil && retEnviCte.protCTe != nil && retEnviCte.protCTe.infProt != nil && retEnviCte.protCTe.infProt.nProt != "" {
		protocolo = retEnviCte.protCTe.infProt.nProt
	}

	arquivoSalvar := filepath.Join(caminhoXml, strings.Builder{}.WriteString(protocolo+"-rec.xml").String())

	return utils.ClasseParaArquivoXml(retEnviCte, arquivoSalvar)
}

// IsAutorizado verifica se o CTe está autorizado
func IsAutorizado(retConsSitCTe *retCTeOS) bool {
	return retConsSitCTe != nil && retConsSitCTe.protCTe != nil && retConsSitCTe.protCTe.infProt.cStat == int(constantes.Autorizado)
}

// IsCancelada verifica se o CTe está cancelado
func IsCancelada(retConsSitCTe *retCTeOS) bool {
	return retConsSitCTe != nil && retConsSitCTe.protCTe != nil && retConsSitCTe.protCTe.infProt.cStat == int(constantes.Cancelada)
}

// IsDenegada verifica se o CTe está denegado
func IsDenegada(retConsSitCTe *retCTeOS) bool {
	return retConsSitCTe != nil && retConsSitCTe.protCTe != nil && (retConsSitCTe.protCTe.infProt.cStat == int(constantes.Denegada) ||
		retConsSitCTe.protCTe.infProt.cStat == int(constants.Denegado205) ||
		retConsSitCTe.protCTe.infProt.cStat == int(constants.DenegadoEmitente))
}

// IsRejeicao verifica se o CTe foi rejeitado
func IsRejeicao(retConsSitCTe *retCTeOS) bool {
	return retConsSitCTe != nil && !IsAutorizado(retConsSitCTe) && !IsCancelada(retConsSitCTe) && !IsDenegada(retConsSitCTe)
}

// IsRejeicao999 verifica se o CTe foi rejeitado com código 999
func IsRejeicao999(retConsSitCTe *retCTeOS) bool {
	return retConsSitCTe != nil && !IsAutorizado(retConsSitCTe) && !IsCancelada(retConsSitCTe) && !IsDenegada(retConsSitCTe) && retConsSitCTe.cStat == 999
}
