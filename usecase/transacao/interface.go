package transacao

import (
	"database/sql"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
)

type Reader interface {
	GetLatestByCliente(db *sql.DB, c *entity.Cliente, quantidade int) ([]*entity.Transacao, error)
}

type Writer interface {
	Create(tx *sql.Tx, clienteId int, t *entity.Transacao) (*entity.Transacao, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUltimasTransacoes(c *entity.Cliente, quantidade int) ([]*entity.Transacao, error)
	CreateTransacao(int, *entity.Transacao) (int, error)
}
