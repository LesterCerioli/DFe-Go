package flags

import "encoding/xml"

// CTeTipoEvento representa os tipos de eventos de CT-e.
type CTeTipoEvento int

const (
	// Evento: Empresa Emitente
	CartaCorrecao                    CTeTipoEvento = 110110
	Cancelamento                     CTeTipoEvento = 110111
	EPEC                             CTeTipoEvento = 110113
	RegistrosdoMultimodal            CTeTipoEvento = 110160
	InformacoesdaGTV                 CTeTipoEvento = 110170
	ComprovantedeEntrega             CTeTipoEvento = 110180
	CancelamentodoComprovantedeEntrega CTeTipoEvento = 110181

	// Evento: Fisco
	RegistrodePassagem               CTeTipoEvento = 310620
	RegistrodePassagemAutomatico      CTeTipoEvento = 510620
	MDFeAutorizado                   CTeTipoEvento = 310610
	MDFeCancelado                    CTeTipoEvento = 310611

	// Evento: Fisco do Emitente
	AutorizadoCTecomplementar         CTeTipoEvento = 240130
	CanceladoCTecomplementar          CTeTipoEvento = 240131
	CTedeSubstituicao                 CTeTipoEvento = 240140
	CTedeAnulacao                     CTeTipoEvento = 240150
	LiberacaodeEPEC                   CTeTipoEvento = 240160
	LiberacaoPrazoCancelamento        CTeTipoEvento = 240170

	// Evento: Ambiente Nacional
	AutorizadoRedespacho              CTeTipoEvento = 440130
	AutorizadoRedespachointermediario CTeTipoEvento = 440140
	AutorizadoSubcontratacao          CTeTipoEvento = 440150
	AutorizadoServicoVinculadoMultimodal CTeTipoEvento = 440160

	// Evento: Tomador
	Desacordo                         CTeTipoEvento = 610110
)

// MarshalXML permite a serialização de CTeTipoEvento como strings XML específicas.
func (e CTeTipoEvento) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: e.String()}, nil
}

// String converte o valor do enum em string para ser usada em XML.
func (e CTeTipoEvento) String() string {
	switch e {
	case CartaCorrecao:
		return "110110"
	case Cancelamento:
		return "110111"
	case EPEC:
		return "110113"
	case RegistrosdoMultimodal:
		return "110160"
	case InformacoesdaGTV:
		return "110170"
	case ComprovantedeEntrega:
		return "110180"
	case CancelamentodoComprovantedeEntrega:
		return "110181"
	case RegistrodePassagem:
		return "310620"
	case RegistrodePassagemAutomatico:
		return "510620"
	case MDFeAutorizado:
		return "310610"
	case MDFeCancelado:
		return "310611"
	case AutorizadoCTecomplementar:
		return "240130"
	case CanceladoCTecomplementar:
		return "240131"
	case CTedeSubstituicao:
		return "240140"
	case CTedeAnulacao:
		return "240150"
	case LiberacaodeEPEC:
		return "240160"
	case LiberacaoPrazoCancelamento:
		return "240170"
	case AutorizadoRedespacho:
		return "440130"
	case AutorizadoRedespachointermediario:
		return "440140"
	case AutorizadoSubcontratacao:
		return "440150"
	case AutorizadoServicoVinculadoMultimodal:
		return "440160"
	case Desacordo:
		return "610110"
	default:
		return ""
	}
}
