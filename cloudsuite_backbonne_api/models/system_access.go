package models

import (
	"encoding/xml"
	"strconv"
)

type SystemAccess struct {
	SystemId int    `xml:"SystemId"`
	Name     string `xml:"Name"`
	Address  string `xml:"Address"`
	IsProper bool   `xml:"IsProper"`
}

func NewSystemAccess(data []byte) (SystemAccess, error) {
	var access SystemAccess
	if err := xml.Unmarshal(data, &access); err != nil {
		return SystemAccess{}, err
	}

	if access.IsProperStr == "true" {
		access.IsProper = true
	} else {
		access.IsProper = false
	}

	return access, nil
}

func (s *SystemAccess) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var elem struct {
		SystemId string `xml:"SystemId"`
		Name     string `xml:"Name"`
		Address  string `xml:"Address"`
		IsProper string `xml:"IsProper"`
	}

	if err := d.DecodeElement(&elem, &start); err != nil {
		return err
	}

	s.Name = elem.Name
	s.Address = elem.Address
	s.IsProper = (elem.IsProper == "true")

	if id, err := strconv.Atoi(elem.SystemId); err == nil {
		s.SystemId = id
	}

	return nil
}
