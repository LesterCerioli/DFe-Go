package event

import (
	"encoding/xml"
)

// RetEnvEvento representa o retorno de um evento
type RetEnvEvento struct {
	Version    string          `xml:"versao,attr"`                  // HR02 - Versão do leiaute
	IdLote     int64           `xml:"idLote"`                       // HR03 - Identificador de controle do Lote de envio do Evento
	TpAmb      TipoAmbiente    `xml:"tpAmb"`                        // HR04 - Identificação do Ambiente
	VerAplic   string          `xml:"verAplic"`                     // HR05 - Versão da aplicação que processou o evento
	COrgao     Estado          `xml:"cOrgao"`                       // HR06 - Código da UF que registrou o Evento
	CStat      int             `xml:"cStat"`                        // HR07 - Código do status da resposta
	XMotivo    string          `xml:"xMotivo"`                      // HR08 - Descrição do status da resposta
	RetEvento   []RetEventoCTe  `xml:"retEvento"`                    // HR09 - TAG de grupo do resultado do processamento do Evento
}

// RetEventoCTe representa a estrutura de retorno de evento CTe (defina conforme necessário)
type RetEventoCTe struct {
	// Campos para retEventoCTe devem ser definidos aqui
}
