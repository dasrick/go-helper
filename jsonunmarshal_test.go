package gohelper

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testUnmarshalTarget struct {
	WhateverID string `json:"Id" validate:"required"`
}

var validJSONString = `{"Id":"this-property-is-required"}`
var invalidJSONString = ``

func TestJSONUnmarshalByte(t *testing.T) {
	validJSONbyte := []byte(validJSONString)
	invalidJSONbyte := []byte(invalidJSONString)
	var tests = []struct {
		value  []byte
		target interface{}
		error  error
	}{
		{
			value:  validJSONbyte,
			target: testUnmarshalTarget{},
			error:  nil,
		},
		{
			value:  invalidJSONbyte,
			target: testUnmarshalTarget{},
			error:  &json.SyntaxError{},
		},
	}
	for _, test := range tests {
		err := JSONUnmarshalByte(test.value, &test.target)
		assert.IsType(t, test.error, err)
	}
}

func TestJSONUnmarshalString(t *testing.T) {
	var tests = []struct {
		value  string
		target interface{}
		error  error
	}{
		{
			value:  validJSONString,
			target: testUnmarshalTarget{},
			error:  nil,
		},
		{
			value:  invalidJSONString,
			target: testUnmarshalTarget{},
			error:  &json.SyntaxError{},
		},
	}
	for _, test := range tests {
		err := JSONUnmarshalString(test.value, &test.target)
		assert.IsType(t, test.error, err)
	}
}

func TestJSONUnmarshalInterface(t *testing.T) {
	validStruct := testUnmarshalTarget{}
	_ = json.Unmarshal([]byte(validJSONString), &validStruct)
	invalidStruct := testUnmarshalTarget{}
	_ = json.Unmarshal([]byte(invalidJSONString), &invalidStruct)
	var tests = []struct {
		value  interface{}
		target interface{}
		error  error
	}{
		{
			value:  make(chan int),
			target: &testUnmarshalTarget{},
			error:  &json.UnsupportedTypeError{},
		},
		{
			value:  validStruct,
			target: &testUnmarshalTarget{},
			error:  nil,
		},
	}
	for _, test := range tests {
		err := JSONUnmarshalInterface(test.value, &test.target)
		assert.IsType(t, test.error, err)
	}
}

//func TestJSONUnmarshalByteAndValidate(t *testing.T) {
//	//validJSONbyte := []byte(validJSONString)
//	invalidJSONbyte := []byte(invalidJSONString)
//	var tests = []struct {
//		value  []byte
//		target interface{}
//		error  error
//	}{
//		{
//			value:  invalidJSONbyte,
//			target: &testUnmarshalTarget{},
//			error:  &json.SyntaxError{},
//		},
//		//{
//		//	value:  validJSONbyte,
//		//	target: &testUnmarshalTarget{},
//		//	error:  nil,
//		//},
//	}
//	for _, test := range tests {
//		err := JSONUnmarshalByteAndValidate(test.value, &test.target)
//		assert.IsType(t, test.error, err)
//	}
//}
