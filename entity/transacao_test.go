package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransacao(t *testing.T) {
	t1, err := NewTransacao(0, "c", "t1")
	assert.Nil(t, t1)
	assert.ErrorIs(t, err, ErrValorZeroOuNegativo)

	t2, err := NewTransacao(1000, "j", "t2")
	assert.Nil(t, t2)
	assert.ErrorIs(t, err, ErrTipoInvalido)

	t3, err := NewTransacao(2000, "", "t3")
	assert.Nil(t, t3)
	assert.ErrorIs(t, err, ErrTipoInvalido)

	t4, err := NewTransacao(123, "d", "descricao longa maior do que 10 caracteres")
	assert.Nil(t, t4)
	assert.ErrorIs(t, err, ErrDescricaoMuitoLonga)

	t5, err := NewTransacao(100, "c", "t4")
	assert.NotNil(t, t5)
	assert.NoError(t, err)
}
