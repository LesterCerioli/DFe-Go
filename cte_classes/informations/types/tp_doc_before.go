package types

// tpDocAnterior representa os tipos de documentos fiscais anteriores.
type tpDocAnterior string

const (
	// CTRC representa o tipo de documento CTRC.
	CTRC tpDocAnterior = "00" // 00 - CTRC
	// CTAC representa o tipo de documento CTAC.
	CTAC tpDocAnterior = "01" // 01 - CTAC
	// ACT representa o tipo de documento ACT.
	ACT tpDocAnterior = "02" // 02 - ACT
	// NFModelo7 representa o tipo de documento NFModelo7.
	NFModelo7 tpDocAnterior = "03" // 03 - NF Modelo 7
	// NFModelo27 representa o tipo de documento NFModelo27.
	NFModelo27 tpDocAnterior = "04" // 04 - NF Modelo 27
	// ConhecimentoAereoNacional representa o tipo de documento de Conhecimento Aéreo Nacional.
	ConhecimentoAereoNacional tpDocAnterior = "05" // 05 - Conhecimento Aéreo Nacional
	// CTMC representa o tipo de documento CTMC.
	CTMC tpDocAnterior = "06" // 06 - CTMC
	// ARTE representa o tipo de documento ARTE.
	ARTE tpDocAnterior = "07" // 07 - ARTE
	// DTA representa o tipo de documento DTA.
	DTA tpDocAnterior = "08" // 08 - DTA
	// ConhecimentoAereoInternacional representa o tipo de documento de Conhecimento Aéreo Internacional.
	ConhecimentoAereoInternacional tpDocAnterior = "09" // 09 - Conhecimento Aéreo Internacional
	// ConhecimentoCartaDePorteInternacional representa o tipo de documento de Conhecimento Carta de Porte Internacional.
	ConhecimentoCartaDePorteInternacional tpDocAnterior = "10" // 10 - Conhecimento Carta de Porte Internacional
	// ConhecimentoAvulso representa o tipo de documento de Conhecimento Avulso.
	ConhecimentoAvulso tpDocAnterior = "11" // 11 - Conhecimento Avulso
	// TIF representa o tipo de documento TIF.
	TIF tpDocAnterior = "12" // 12 - TIF
	// BL representa o tipo de documento BL.
	BL tpDocAnterior = "13" // 13 - BL
	// Outros representa outros tipos de documentos.
	Outros tpDocAnterior = "99" // 99 - Outros
)
