package cliente

import (
	"context"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/amardegan/rinha-de-backend-2024-q1/infrastructure/repository"
)

type Service struct {
	repo *repository.DBConn
}

func NewService(r *repository.DBConn) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ListClientes(ctx context.Context) ([]*entity.Cliente, error) {
	clientes, err := repository.ListClientes(ctx, s.repo.Pool)
	if err != nil {
		return nil, err
	}
	if len(clientes) == 0 {
		return nil, entity.ErrClienteNaoEncontrado
	}
	return clientes, nil
}

func (s *Service) GetClienteById(ctx context.Context, clienteId int) (*entity.Cliente, error) {
	return repository.GetClienteById(ctx, s.repo.Pool, clienteId)
}
