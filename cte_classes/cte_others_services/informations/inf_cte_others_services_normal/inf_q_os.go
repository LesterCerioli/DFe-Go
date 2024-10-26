package inf_cte_others_services_normal

// InfQOs é a estrutura equivalente em Go da classe C# infQOs
type InfQOs struct {
	QCarga         *float64 `xml:"qCarga"`
}

// SetQCarga arredonda e define o valor de QCarga
func (i *InfQOs) SetQCarga(value float64) {
	roundedValue := round(value, 4)
	i.QCarga = &roundedValue
}

// QCargaSpecified verifica se o valor de QCarga está presente
func (i *InfQOs) QCargaSpecified() bool {
	return i.QCarga != nil
}

// round arredonda um número com precisão especificada
func round(value float64, precision int) float64 {
	format := "%." + fmt.Sprintf("%df", precision)
	result, _ := strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return result
}
