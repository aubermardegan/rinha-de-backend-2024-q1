package cliente

import (
	"context"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Reader interface {
	List(context.Context, *pgxpool.Pool) ([]*entity.Cliente, error)
	Get(context.Context, *pgxpool.Pool) (*entity.Cliente, error)
}

type Repository interface {
	Reader
}

type UseCase interface {
	ListClientes(context.Context) ([]*entity.Cliente, error)
	GetClienteById(context.Context, int) (*entity.Cliente, error)
}
