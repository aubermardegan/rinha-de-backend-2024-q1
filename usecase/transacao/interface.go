package transacao

import (
	"database/sql"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
)

type Reader interface {
	GetByClienteAfterId(tx *sql.Tx, clienteId, transacaoId int) ([]*entity.Transacao, error)
	GetLatestByCliente(clienteId, quantidade int) ([]*entity.Transacao, error)
}

type Writer interface {
	Create(tx *sql.Tx, clienteId int, t *entity.Transacao) (*entity.Transacao, error)
}

type TransactionControl interface {
	BeginTran() (*sql.Tx, error)
	CommitTran(*sql.Tx) error
	RollbackTran(*sql.Tx) error
}

type Repository interface {
	Reader
	Writer
	TransactionControl
}

type UseCase interface {
	GetUltimasTransacoes(clienteId, quantidade int) ([]*entity.Transacao, error)
	CreateTransacao(*sql.Tx, *entity.Cliente, *entity.Transacao) (*entity.Transacao, error)
}
