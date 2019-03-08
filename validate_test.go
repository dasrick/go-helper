package gohelper

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

type testStruct struct {
	ContractID string `json:"Id" validate:"required"`
}

func TestHandler(t *testing.T) {
	validJSONbyte := []byte(`{"Id":"this-property-is-required"}`)
	invalidJSONbyte := []byte(`{"useless":"property"}`)
	validStruct := &testStruct{}
	invalidStruct := &testStruct{}
	_ = json.Unmarshal(validJSONbyte, validStruct)
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
