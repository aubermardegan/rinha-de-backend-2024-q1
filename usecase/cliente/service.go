package cliente

import (
	"database/sql"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/amardegan/rinha-de-backend-2024-q1/infrastructure/repository"
)

type Service struct {
	repo *sql.DB
}

func NewService(r *sql.DB) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ListClientes() ([]*entity.Cliente, error) {
	clientes, err := repository.ListClientes(s.repo)
	if err != nil {
		return nil, err
	}
	if len(clientes) == 0 {
		return nil, entity.ErrClienteNaoEncontrado
	}
	return clientes, nil
}

func (s *Service) GetClienteById(clienteId int) (*entity.Cliente, error) {
	return repository.GetClienteById(s.repo, clienteId)
}
