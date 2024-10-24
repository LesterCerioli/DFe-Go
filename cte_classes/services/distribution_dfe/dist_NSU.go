package distribution_dfe

import (
	"encoding/xml"
)

// DistNSU representa o grupo para distribuir DF-e de interesse.
type DistNSU struct {
	XMLName xml.Name `xml:"distNSU"`
	UltNSU  string   `xml:"ultNSU"`
}
