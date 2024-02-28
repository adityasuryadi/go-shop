package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Validation struct {
	validation *validator.Validate
}

func NewValidation(db *gorm.DB) *Validation {
	validation := validator.New()

	// register custom unique validation
	validation.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		// fmt.Println(fl.StructFieldName())
		// // get parameter dari tag struct validate
		table := fl.Param()
		field := strings.ToLower(fl.StructFieldName())
		var total int64
		err := db.Table(table).Where(""+field+" = ?", fl.Field()).Count(&total).Error
		if err != nil {
			fmt.Println(err)
		}
		// // Return true if the count is zero (i.e., the value is unique)
		return total == 0
	})

	validation.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

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
		return fe.Field() + " not avaiable"
	case "min":
		return "minimal " + fe.Param() + " karakter"
	case "max":
		return "max " + fe.Param() + " Kb"
	case "image_validation":
		return "Harus Image"
	case "number":
		return "harus numeric"
	case "eqfield":
		return "field tidak sama dengan " + fe.Param()
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
