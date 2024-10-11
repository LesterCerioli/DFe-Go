package models

// TipoAmbiente representa o tipo de ambiente (produção ou homologação).
type TipoAmbiente string

// Versao representa a versão do serviço.
type Versao string

// ConfigWebService representa a configuração do web service para emissão de NFe e NFC-e.
type ConfigWebService struct {
	UfEmitente    Estado        `json:"ufEmitente"`
	Ambiente      TipoAmbiente  `json:"ambiente"`
	Serie         int16         `json:"serie"`        
	Numeracao     int64         `json:"numeracao"`     
	Versao        Versao        `json:"versao"`
	CaminhoSchemas string        `json:"caminhoSchemas"`
	TimeOut       int           `json:"timeOut"`
}
