package models

import "encoding/json"

type ConfigureDigitalCertificate struct {
	SerieNumber string `json:"serie_number"`
	FileWay     string `json:"file_way"`
	Password    string `json:"password"`
	KeepInCache bool   `json:"keep_in_cache"`
}

func main() {
	// Example usage
	config := ConfigureDigitalCertificate{
		SerieNumber: "123456",
		FileWay:     "/path/to/file",
		Password:    "password",
		KeepInCache: true,
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(config)
	println(string(jsonData))
}
