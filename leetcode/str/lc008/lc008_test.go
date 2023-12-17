package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	testCases := []struct {
		input string
		value int
	}{
		{input: "42", value: 42},
		{input: "   -42", value: -42},
		{input: "4193 with words", value: 4193},
		{input: "words and 987", value: 0},
		{input: "-91283472332", value: -2147483648},
		{input: "   +0 123", value: 0},
		{input: "9223372036854775808", value: 2147483647},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			myAtoi(testCase.input)
			assert := assert.New(t)
			assert.Equal(testCase.value, myAtoi(testCase.input))
		})
	}
}
