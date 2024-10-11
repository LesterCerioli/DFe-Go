package danfe_shared_helper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"

	"github.com/yourusername/fastreport" // Replace with actual import path
	"github.com/yourusername/nfe"        // Replace with actual import path
)

type DanfeSharedHelper struct{}

func (d *DanfeSharedHelper) GenerateDanfeNfceReport(proc nfe.NfeProc, configuracaoDanfeNfce nfe.ConfiguracaoDanfeNfce, cIdToken, csc string, frx []byte, arquivoRelatorio, textoRodape string) (*fastreport.Report, error) {
	relatorio := fastreport.NewReport()

	if arquivoRelatorio != "" {
		if err := relatorio.Load(arquivoRelatorio); err != nil {
			return nil, err
		}
	} else if len(frx) > 0 {
		if err := relatorio.Load(io.NopCloser(bytes.NewReader(frx))); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Erro em DanfeSharedHelper.GenerateDanfeNfceReport: Relatório não encontrado, passe os parametros 'frx' com bytes ou 'arquivoRelatorio' com o caminho do arquivo")
	}

	relatorio.RegisterData([]nfe.NfeProc{proc}, "NFCe", 20)
	relatorio.GetDataSource("NFCe").Enabled = true
	relatorio.SetParameterValue("NfceDetalheVendaNormal", configuracaoDanfeNfce.DetalheVendaNormal)
	relatorio.SetParameterValue("NfceDetalheVendaContigencia", configuracaoDanfeNfce.DetalheVendaContigencia)
	relatorio.SetParameterValue("NfceImprimeDescontoItem", configuracaoDanfeNfce.ImprimeDescontoItem)
	relatorio.SetParameterValue("NfceImprimeFoneEmitente", configuracaoDanfeNfce.ImprimeFoneEmitente)

	var foneEmitente string
	if proc.NFe.InfNFe.Emit.EnderEmit.Fone != "" {
		foneEmitente = proc.NFe.InfNFe.Emit.EnderEmit.Fone
	}

	if len(foneEmitente) == 10 {
		foneEmitente = fmt.Sprintf("%s:(00)0000-0000", foneEmitente)
	} else if len(foneEmitente) == 11 {
		foneEmitente = fmt.Sprintf("%s:(00)00000-0000", foneEmitente)
	}

	relatorio.SetParameterValue("NfceFoneEmitente", foneEmitente)
	relatorio.SetParameterValue("NfceModoImpressao", configuracaoDanfeNfce.ModoImpressao)
	relatorio.SetParameterValue("NfceCancelado", configuracaoDanfeNfce.DocumentoCancelado)
	relatorio.SetParameterValue("NfceLayoutQrCode", configuracaoDanfeNfce.NfceLayoutQrCode)
	relatorio.SetParameterValue("TextoRodape", textoRodape)

	relatorio.FindObject("PgNfce").SetMargins(configuracaoDanfeNfce.MargemEsquerda, configuracaoDanfeNfce.MargemDireita)

	logomarcaEmitDefinida := configuracaoDanfeNfce.Logomarca != nil && len(configuracaoDanfeNfce.Logomarca) > 0
	relatorio.FindObject("rtbEmitLogo").SetVisible(logomarcaEmitDefinida)
	if logomarcaEmitDefinida {
		relatorio.FindObject("poEmitLogo").SetImageData(configuracaoDanfeNfce.Logomarca)
	}

	if proc.NFe.InfNFeSupl.UrlChave == "" {
		relatorio.FindObject("txtUrl").SetText(proc.NFe.InfNFeSupl.ObterUrlConsulta(proc.NFe, configuracaoDanfeNfce.VersaoQrCode))
	} else {
		relatorio.FindObject("txtUrl").SetText(proc.NFe.InfNFeSupl.UrlChave)
	}

	if proc.NFe.InfNFeSupl == nil {
		relatorio.FindObject("bcoQrCode").SetText(proc.NFe.InfNFeSupl.ObterUrlQrCode(proc.NFe, configuracaoDanfeNfce.VersaoQrCode, cIdToken, csc))
		relatorio.FindObject("bcoQrCodeLateral").SetText(proc.NFe.InfNFeSupl.ObterUrlQrCode(proc.NFe, configuracaoDanfeNfce.VersaoQrCode, cIdToken, csc))
	} else {
		relatorio.FindObject("bcoQrCode").SetText(proc.NFe.InfNFeSupl.QrCode)
		relatorio.FindObject("bcoQrCodeLateral").SetText(proc.NFe.InfNFeSupl.QrCode)
	}

	if configuracaoDanfeNfce.SegundaViaContingencia {
		// Handle print settings based on conditions
	}

	return relatorio, nil
}

func (d *DanfeSharedHelper) GenerateDanfeFrEventoReport(proc nfe.NfeProc, procEventoNFe nfe.ProcEventoNFe, configuracaoDanfeNfe nfe.ConfiguracaoDanfeNfe, frx []byte, desenvolvedor, arquivoRelatorio string) (*fastreport.Report, error) {
	relatorio := fastreport.NewReport()

	if arquivoRelatorio != "" {
		if err := relatorio.Load(arquivoRelatorio); err != nil {
			return nil, err
		}
	} else if len(frx) > 0 {
		if err := relatorio.Load(io.NopCloser(bytes.NewReader(frx))); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Erro em DanfeSharedHelper.GenerateDanfeFrNfeReport: Relatório não encontrado, passe os parametros 'frx' com bytes ou 'arquivoRelatorio' com o caminho do arquivo")
	}

	relatorio.RegisterData([]nfe.NfeProc{proc}, "NFe", 20)
	relatorio.RegisterData([]nfe.ProcEventoNFe{procEventoNFe}, "procEventoNFe", 20)
	relatorio.GetDataSource("NFe").Enabled = true
	relatorio.GetDataSource("procEventoNFe").Enabled = true
	relatorio.SetParameterValue("desenvolvedor", desenvolvedor)

	return relatorio, nil
}

func (d *DanfeSharedHelper) GenerateDanfeFrNfeReport(proc nfe.NfeProc, configuracaoDanfeNfe nfe.ConfiguracaoDanfeNfe, frx []byte, desenvolvedor, arquivoRelatorio string) (*fastreport.Report, error) {
	relatorio := fastreport.NewReport()

	if arquivoRelatorio != "" {
		if err := relatorio.Load(arquivoRelatorio); err != nil {
			return nil, err
		}
	} else if len(frx) > 0 {
		if err := relatorio.Load(io.NopCloser(bytes.NewReader(frx))); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Erro em DanfeSharedHelper.GenerateDanfeFrNfeReport: Relatório não encontrado, passe os parametros 'frx' com bytes ou 'arquivoRelatorio' com o caminho do arquivo")
	}

	relatorio.RegisterData([]nfe.NfeProc{proc}, "NFe", 20)
	relatorio.GetDataSource("NFe").Enabled = true

	var mensagem, resumoCanhoto, contingenciaDescricao, contingenciaValor, consultaAutenticidade string
	consultaAutenticidade = "Consulta de autenticidade no portal nacional da NF-e\nwww.nfe.fazenda.gov.br/portal ou no site da Sefaz autorizadora"

	if configuracaoDanfeNfe.ExibirResumoCanhoto {
		if configuracaoDanfeNfe.ResumoCanhoto == "" {
			resumoCanhoto = fmt.Sprintf("Emissão: %s Dest/Reme: %s Valor Total: %.2f", proc.NFe.InfNFe.Ide.DhEmi.Format("02/01/2006"), proc.NFe.InfNFe.Dest.XNome, proc.NFe.InfNFe.Total.ICMSTot.VNF)
		} else {
			resumoCanhoto = configuracaoDanfeNfe.ResumoCanhoto
		}
	}

	if proc.NFe.InfNFe.Ide.TpAmb == nfe.TipoAmbienteHomologacao {
		if proc.NFe.InfNFe.Ide.TpEmis == nfe.TipoEmissaoTeSCAN || proc.NFe.InfNFe.Ide.TpEmis == nfe.TipoEmissaoTeEPEC || proc.NFe.InfNFe.Ide.TpEmis == nfe.TipoEmissaoTeFSDA || proc.NFe.InfNFe.Ide.TpEmis == nfe.TipoEmissaoTeFSIA {
			if proc.ProtNFe != nil && proc.ProtNFe.InfProt != nil && (proc.ProtNFe.InfProt.CStat == 101 || proc.ProtNFe.InfProt.CStat == 135 || proc.ProtNFe.InfProt.CStat == 151 || proc.ProtNFe.InfProt.CStat == 155) {
				mensagem = "NFe sem Valor Fiscal - HOMOLOGAÇÃO\nNFe em Contingência - CANCELADA"
			} else {
				mensagem = "NFe sem Valor Fiscal - HOMOLOGAÇÃO\nNFe em Contingência"
			}
		} else {
			mensagem = "NFe sem Valor Fiscal - HOMOLOGAÇÃO"
		}
	} else {
		if configuracaoDanfeNfe.DocumentoCancelado || (proc.ProtNFe != nil && proc.ProtNFe.InfProt != nil && proc.ProtNFe.InfProt.NProt != "" && (proc.ProtNFe.InfProt.CStat == 101 || proc.ProtNFe.InfProt.CStat == 135 || proc.ProtNFe.InfProt.CStat == 151 || proc.ProtNFe.InfProt.CStat == 155)) {
			mensagem = "NFe Cancelada"
		} else if proc.ProtNFe != nil && proc.ProtNFe.InfProt != nil && (proc.ProtNFe.InfProt.CStat == 110 || proc.ProtNFe.InfProt.CStat == 301 || proc.ProtNFe.InfProt.CStat == 302 || proc.ProtNFe.InfProt.CStat == 303) {
			mensagem = "NFe denegada pelo Fisco"
		} else if proc.ProtNFe != nil && proc.ProtNFe.InfProt != nil && proc.ProtNFe.InfProt.NProt == "" {
			mensagem = "NFe sem Autorização de Uso da SEFAZ"
		}
	}

	switch proc.NFe.InfNFe.Ide.TpEmis {
	case nfe.TipoEmissaoTeNormal, nfe.TipoEmissaoTeSCAN, nfe.TipoEmissaoTeSVCAN, nfe.TipoEmissaoTeSVCRS:
		contingenciaDescricao = "PROTOCOLO DE AUTORIZAÇÃO DE USO"
		if proc.ProtNFe == nil || proc.ProtNFe.InfProt == nil || proc.ProtNFe.InfProt.NProt == "" {
			contingenciaValor = "NFe sem Autorização de Uso da SEFAZ"
		} else {
			contingenciaValor = fmt.Sprintf("%s - %s", proc.ProtNFe.InfProt.NProt, proc.ProtNFe.InfProt.DhRecbto.Format("02/01/2006 15:04:05"))
		}
		if configuracaoDanfeNfe.DocumentoCancelado || (proc.ProtNFe != nil && proc.ProtNFe.InfProt != nil && (proc.ProtNFe.InfProt.CStat == 101 || proc.ProtNFe.InfProt.CStat == 151 || proc.ProtNFe.InfProt.CStat == 155)) {
			contingenciaDescricao = "PROTOCOLO DE HOMOLOGAÇÃO DO CANCELAMENTO"
		} else if proc.ProtNFe != nil && proc.ProtNFe.InfProt != nil && (proc.ProtNFe.InfProt.CStat == 110 || proc.ProtNFe.InfProt.CStat == 301 || proc.ProtNFe.InfProt.CStat == 302 || proc.ProtNFe.InfProt.CStat == 303) {
			contingenciaDescricao = "PROTOCOLO DE DENEGAÇÃO DE USO"
		}
	case nfe.TipoEmissaoTeFSIA, nfe.TipoEmissaoTeEPEC, nfe.TipoEmissaoTeFSDA:
		contingenciaDescricao = "DADOS DA NF-E"
		contingenciaValor = regexp.MustCompile(".{4}").ReplaceAllString(configuracaoDanfeNfe.ChaveContingencia, "$0 ")
		consultaAutenticidade = ""
	default:
		contingenciaValor = fmt.Sprintf("%s - %s", proc.ProtNFe.InfProt.NProt, proc.ProtNFe.InfProt.DhRecbto.Format("02/01/2006 15:04:05"))
	}

	relatorio.SetParameterValue("ResumoCanhoto", resumoCanhoto)
	relatorio.SetParameterValue("Mensagem", mensagem)
	relatorio.SetParameterValue("ConsultaAutenticidade", consultaAutenticidade)
	relatorio.SetParameterValue("ContingenciaDescricao", contingenciaDescricao)
	relatorio.SetParameterValue("ContingenciaValor", contingenciaValor)
	relatorio.SetParameterValue("ContingenciaID", configuracaoDanfeNfe.ChaveContingencia)
	relatorio.SetParameterValue("DuasLinhas", configuracaoDanfeNfe.DuasLinhas)
	relatorio.SetParameterValue("Desenvolvedor", desenvolvedor)
	relatorio.SetParameterValue("QuebrarLinhasObservacao", configuracaoDanfeNfe.QuebrarLinhasObservacao)
	relatorio.SetParameterValue("ImprimirISSQN", configuracaoDanfeNfe.ImprimirISSQN)
	relatorio.SetParameterValue("ImprimirDescPorc", configuracaoDanfeNfe.ImprimirDescPorc)
	relatorio.SetParameterValue("ImprimirTotalLiquido", configuracaoDanfeNfe.ImprimirTotalLiquido)
	relatorio.SetParameterValue("ImprimirUnidQtdeValor", configuracaoDanfeNfe.ImprimirUnidQtdeValor)
	relatorio.SetParameterValue("ExibeCampoFatura", configuracaoDanfeNfe.ExibeCampoFatura)
	relatorio.SetParameterValue("Logo", configuracaoDanfeNfe.Logomarca)
	relatorio.SetParameterValue("ExibirTotalTributos", configuracaoDanfeNfe.ExibirTotalTributos)
	relatorio.SetParameterValue("ExibeRetencoes", configuracaoDanfeNfe.ExibeRetencoes)
	relatorio.SetParameterValue("DecimaisValorUnitario", configuracaoDanfeNfe.DecimaisValorUnitario)
	relatorio.SetParameterValue("DecimaisQuantidadeItem", configuracaoDanfeNfe.DecimaisQuantidadeItem)
	relatorio.SetParameterValue("DataHoraImpressao", configuracaoDanfeNfe.DataHoraImpressao)

	return relatorio, nil
}
