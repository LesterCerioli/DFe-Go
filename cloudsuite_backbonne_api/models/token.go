package models

import (
	"github.com/go-playground/validator/v10"
)

type Token struct {
	Value     string `json:"value" validate:"required"`
	IPAddress string `json:"ip_address" validate:"required"`
}

func (t *Token) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
