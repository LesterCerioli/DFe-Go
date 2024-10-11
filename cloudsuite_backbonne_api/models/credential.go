package models

import (
	"github.com/go-playground/validator/v10"
)

type Credential struct {
	Login     string `json:"login" validate:"required"`
	Password  string `json:"password" validate:"required"`
	IPAddress string `json:"ip_address" validate:"required,ip"`
}

func ValidateCredential(cred Credential) error {
	validate := validator.New()
	return validate.Struct(cred)
}
