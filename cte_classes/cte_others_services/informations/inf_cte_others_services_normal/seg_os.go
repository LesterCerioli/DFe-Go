package inf_cte_others_services_normal

// SegOs representa a estrutura equivalente em Go da classe C# segOs.
type SegOs struct {
	RespSeg   RespSeg `xml:"respSeg"`   // Campo para resposta do seguro
	XSeg      string  `xml:"xSeg"`      // Campo para descrição do seguro
	NApol     string  `xml:"nApol"`     // Campo para número da apólice
}

// RespSeg deve ser definida conforme a classe C# original.
type RespSeg struct {
	// Definição dos campos da classe respSeg (a ser ajustada com base na classe C# original)
}
