package gohelper

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

// ValidateStruct ...
func ValidateStruct(target interface{}) error {
	if validate == nil {
		validate = validator.New()
	}
	return validate.Struct(target)
}
