package entity

import (
	"errors"
	"sync"
)

type Cliente struct {
	Id     int
	Limite int
	Saldo  Saldo
}

type Saldo struct {
	Valor                      int
	UltimoIdTransacaoConferido int
}

func NewCliente(id, limite, saldoInicial int) (*Cliente, error) {

	c := &Cliente{
		Id:     id,
		Limite: limite,
		Saldo: Saldo{
			Valor:                      saldoInicial,
			UltimoIdTransacaoConferido: 0,
		},
	}

	err := c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Cliente) Validate() error {

	if c.Id <= 0 {
		return ErrIdClienteInvalido
	}

	if c.Limite < 0 {
		return ErrLimiteInvalido
	}

	if c.Saldo.Valor < (c.Limite * -1) {
		return ErrSaldoInferiorAoLimite
	}

	return nil
}

func (c *Cliente) SemLimiteDisponivel(t *Transacao) bool {
	return t.Tipo == TransacaoDebito && (c.Saldo.Valor-t.Valor) < (c.Limite*-1)
}

var (
	ErrIdClienteInvalido     = errors.New("o id do cliente deve ser um número positivo")
	ErrLimiteInvalido        = errors.New("o limite não pode ser negativo")
	ErrSaldoInferiorAoLimite = errors.New("o saldo não pode extrapolar o limite")
	ErrClienteNaoEncontrado  = errors.New("cliente não encontrado")
)

type BufferClientes struct {
	sync.RWMutex
	clientes map[int]*Cliente
}

func InitBufferClientes() BufferClientes {
	return BufferClientes{
		clientes: map[int]*Cliente{},
	}
}

func (b *BufferClientes) AddCliente(c *Cliente) {
	b.Lock()
	_, ok := b.clientes[c.Id]
	if !ok {
		b.clientes[c.Id] = c
	}
	b.Unlock()
}

func (b *BufferClientes) GetCliente(id int) (c *Cliente, ok bool) {
	b.RLock()
	c, ok = b.clientes[id]
	b.RUnlock()
	return
}
