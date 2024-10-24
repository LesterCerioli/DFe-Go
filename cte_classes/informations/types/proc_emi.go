package types

// procEmi representa os processos de emissão da NF-e.
type procEmi string

const (
	// AplicativoContribuinte representa a emissão de NF-e com aplicativo do contribuinte.
	AplicativoContribuinte procEmi = "0" // 0 - emissão de NF-e com aplicativo do contribuinte
	// AvulsaFisco representa a emissão de NF-e avulsa pelo Fisco.
	AvulsaFisco procEmi = "1" // 1 - emissão de NF-e avulsa pelo Fisco
	// AvulsaContribuinte representa a emissão de NF-e avulsa pelo contribuinte com seu certificado digital, através do site do Fisco.
	AvulsaContribuinte procEmi = "2" // 2 - emissão de NF-e avulsa pelo contribuinte
	// ContribuinteAplicativoFisco representa a emissão de NF-e pelo contribuinte com aplicativo fornecido pelo Fisco.
	ContribuinteAplicativoFisco procEmi = "3" // 3 - emissão de NF-e pelo contribuinte com aplicativo fornecido pelo Fisco
)
