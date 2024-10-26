package sender

// InfNF representa a estrutura da informação da Nota Fiscal.
type InfNF struct {
	LocRet LocRet   // Local de Retirada

	NRoma  string   // Número da Roma
	NPed   string   // Número do Pedido
	Mod    int      // Modelo
	Serie  string   // Série
	NDoc   string   // Número do Documento
	DEmi   string   // Data de Emissão
	VBC    float64  // Valor da Base de Cálculo
	VICMS  float64  // Valor do ICMS
	VBCST  float64  // Valor da Base de Cálculo do ICMS ST
	VST    float64  // Valor do ICMS ST
	VProd  float64  // Valor Total dos Produtos
	VNF    float64  // Valor Total da Nota Fiscal
	NCFOP  int      // Código Fiscal de Operações e Prestações
	NPeso  float64  // Peso Total
	PIN    string   // PIN
}
