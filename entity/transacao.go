package entity

import "errors"

type Transacao struct {
	Valor     int    `json:"valor"`
	Tipo      byte   `json:"tipo"`
	Descricao string `json:"descricao"`
}

func NewTransacao(valor int, tipo, descricao string) (*Transacao, error) {

	var byteTipo byte
	if len(tipo) > 0 {
		byteTipo = tipo[0]
	}

	t := &Transacao{
		Valor:     valor,
		Tipo:      byteTipo,
		Descricao: descricao,
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

	if len(t.Descricao) == 0 {
		return ErrCampoVazio
	}

	if t.Tipo != tipoCredito && t.Tipo != tipoDebito {
		return ErrTipoInvalido
	}
	return nil
}

const tipoCredito = 'c'
const tipoDebito = 'd'

var (
	ErrValorZeroOuNegativo = errors.New("o valor não pode ser negativo")
	ErrCampoVazio          = errors.New("o campo não pode ser vazio")
	ErrTipoInvalido        = errors.New("tipo inválido")
)
