package main

import (
	"testing"

	//"github.com/hoisie/mustache"
	"github.com/stretchr/testify/assert"
)

func TestDivideString(t *testing.T) {
	s := divideString("")
	assert.Equal(t, s, "", "El string esta vacio")
}

func TestAsign(t *testing.T) {
	var cases = []struct {
		Type   string
		Length int
		Value  string
	}{
		{"NNN", 987654321, "10"},
		{"", 0, ""},
		{"", 10, "JDD"},
		{"0004", 5000000, ""},
	}
	for _, testData := range cases {
		lt1 := len(testData.Type)
		if lt1 != 2 {
			m := (lt1 != 2)
			assert.Equal(t, m, "El tamaño de el primer parametro debe ser 2")
		}
		lt2 := testData.Length
		if lt2 > 99 {
			m := lt2 > 99
			assert.Equal(t, m, "El segundo parametro debe ser menor a 99")
		}
		lt3 := len(testData.Value)
		if lt3 == 0 {
			m := (lt3 == 0)
			assert.Equal(t, m, "El tercer parametro esta vacio")
		}
	}
}

/*
func TestParser(t *testing.T) {
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
		r, err := mustache.ParseString([]byte(testData.Input))
		// acá agregar chequeos propios del test por ejemplo:
		assert.Equal(t, err == nil, testData.Success)
	}
}*/
