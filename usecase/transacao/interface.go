package transacao

import (
	"context"
	"database/sql"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Reader interface {
	GetLatestByCliente(ctx context.Context, db *pgxpool.Pool, c *entity.Cliente, quantidade int) ([]*entity.Transacao, error)
}

type Writer interface {
	Create(tx *sql.Tx, clienteId int, t *entity.Transacao) (*entity.Transacao, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUltimasTransacoes(ctx context.Context, c *entity.Cliente, quantidade int) ([]*entity.Transacao, error)
	CreateTransacao(context.Context, int, *entity.Transacao) (int, error)
}
