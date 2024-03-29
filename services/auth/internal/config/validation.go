package config

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validation struct {
	validation *validator.Validate
}

func NewValidation() *Validation {
	validation := validator.New()
	return &Validation{
		validation: validation,
	}
}

func removeFirstNameSpace(namespace string) string {
	s := strings.Split(namespace, ".")
	if len(s) > 1 {
		arr := make([]string, 0, len(s))
		for i := 1; i < len(s); i++ {
			arr = append(arr, s[i])
		}
		result := strings.Join([]string(arr), ".")
		return result
	}
	return namespace
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field " + fe.Field() + " tidak boleh kosong"
	case "lte":
		return "harus lebih kecil dari " + fe.Param()
	case "gtenow":
		return "harus lebih besar dari tanggal hari ini"
	case "gte":
		return "harus lebih besar dari " + fe.Param()
	case "email":
		return "format email salah"
	case "unique":
		return "data exist"
	case "min":
		return "minimal " + fe.Param() + " karakter"
	case "max":
		return "max " + fe.Param() + " Kb"
	case "image_validation":
		return "Harus Image"
	case "number":
		return "harus numeric"
	}
	return "Unknown error"
}

func (v *Validation) ValidateRequest(request interface{}) error {
	err := v.validation.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

func (v *Validation) ErrorJson(err error) interface{} {

	validationErrors := err.(validator.ValidationErrors)
	out := make(map[string][]string, len(validationErrors))
	for _, fieldError := range validationErrors {
		out[removeFirstNameSpace(fieldError.Namespace())] = append(out[removeFirstNameSpace(fieldError.Namespace())], GetErrorMsg(fieldError))
	}
	return out
}
