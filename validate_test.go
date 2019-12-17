package gohelper

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type testValidationStruct struct {
	WahteverID string `json:"Id" validate:"required"`
}

func TestHandler(t *testing.T) {
	validJSONbyte := []byte(`{"Id":"this-property-is-required"}`)
	invalidJSONbyte := []byte(`{"useless":"property"}`)
	validStruct := &testValidationStruct{}
	invalidStruct := &testValidationStruct{}
	_ = json.Unmarshal(validJSONbyte, &validStruct)
	_ = json.Unmarshal(invalidJSONbyte, invalidStruct)

	var tests = []struct {
		request interface{}
		error   error
	}{
		{
			request: validStruct,
			error:   nil,
		},
		{
			request: invalidStruct,
			error:   validator.ValidationErrors{},
		},
	}

	for _, test := range tests {
		err := ValidateStruct(test.request)
		assert.IsType(t, test.error, err)
	}
}
