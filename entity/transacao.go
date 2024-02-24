package entity

import (
	"errors"
	"time"
)

const TransacaoCredito = "c"
const TransacaoDebito = "d"

var (
	ErrValorZeroOuNegativo    = errors.New("o valor não pode ser negativo")
	ErrDescricaoVazia         = errors.New("a descrição não pode ser vazia")
	ErrDescricaoMuitoLonga    = errors.New("a descrição não pode exceder 10 caracteres")
	ErrCampoVazio             = errors.New("o campo não pode ser vazio")
	ErrTipoInvalido           = errors.New("tipo inválido")
	ErrTransacaoNaoEncontrada = errors.New("nenhuma transacao encontrada")
	ErrSemLimiteParaTransacao = errors.New("sem limite disponivel para concluir a transacao")
)

type Transacao struct {
	Id          int
	Valor       int       `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

func NewTransacao(valor int, tipo, descricao string) (*Transacao, error) {

	t := &Transacao{
		Valor:       valor,
		Tipo:        tipo,
		Descricao:   descricao,
		RealizadaEm: time.Now().UTC(),
	}

	err := t.Validate()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Transacao) Validate() error {

	if t.Valor <= 0 {
		return ErrValorZeroOuNegativo
	}

	if t.Tipo != TransacaoCredito && t.Tipo != TransacaoDebito {
		return ErrTipoInvalido
	}

	if len(t.Descricao) == 0 {
		return ErrCampoVazio
	}

	if len(t.Descricao) > 10 {
		return ErrDescricaoMuitoLonga
	}

	return nil
}
