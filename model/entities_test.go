package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResult(t *testing.T) {
	t_ := "XSC"
	l_ := 04
	v_ := "ÑASJ"
	NewResult(t_, l_, v_)
	lt1 := len(t_)
	if lt1 != 2 {
		m := (lt1 != 2)
		assert.Equal(t, m, "El tamaño de el primer parametro debe ser 2")
	}
	lt2 := l_
	if lt2 > 99 {
		m := lt2 > 99
		assert.Equal(t, m, "El segundo parametro debe ser menor a 99")
	}
	lt3 := len(v_)
	if lt3 == 0 {
		m := (lt3 == 0)
		assert.Equal(t, m, "El tercer parametro esta vacio")
	}
}
