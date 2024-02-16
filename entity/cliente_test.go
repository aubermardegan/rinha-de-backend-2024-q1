package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCliente(t *testing.T) {
	c1, err := NewCliente(0, 0, 0)
	assert.Nil(t, c1)
	assert.ErrorIs(t, err, ErrIdClienteInvalido)

	_, err = NewCliente(-1, 0, 0)
	assert.ErrorIs(t, err, ErrIdClienteInvalido)

	_, err = NewCliente(1, -1, 0)
	assert.ErrorIs(t, err, ErrLimiteInvalido)

	_, err = NewCliente(2, 0, -1)
	assert.ErrorIs(t, err, ErrSaldoInferiorAoLimite)

	_, err = NewCliente(3, 1000, -1001)
	assert.ErrorIs(t, err, ErrSaldoInferiorAoLimite)

	c2, err := NewCliente(4, 1000, -1000)
	assert.NotNil(t, c2)
	assert.NoError(t, err)
}
