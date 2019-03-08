package gohelper

import "gopkg.in/go-playground/validator.v9"

var validate *validator.Validate

// ValidateStruct ...
func ValidateStruct(target interface{}) error {
	if validate == nil {
		validate = validator.New()
	}
	return validate.Struct(target)
}
