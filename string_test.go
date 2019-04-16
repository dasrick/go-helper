package gohelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringDeduplicate(t *testing.T) {
	var tests = []struct {
		request []string
		result  []string
	}{
		{
			request: make([]string, 0),
			result:  []string{},
		},
		{
			request: []string{"abc", "def", "abc", "abc"},
			result:  []string{"abc", "def"},
		},
	}
	for _, test := range tests {
		result := StringDeduplicate(test.request)
		assert.Equal(t, test.result, result)
	}
}
