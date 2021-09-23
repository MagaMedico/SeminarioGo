package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideString(t *testing.T) {
	s := divideString("")
	assert.Equal(t, s, "", "El string esta vacio")
}

func TestAsign(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}
	for _, testData := range cases {

		a := divideString(testData.Input)

		test := asign(a) //Convert to string

		assert.Equal(t, testData.Type, test.Type)
		assert.Equal(t, testData.Value, test.Value)
		assert.Equal(t, testData.Length, test.Length)
	}
}
