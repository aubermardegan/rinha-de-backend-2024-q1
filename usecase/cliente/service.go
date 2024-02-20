package cliente

import "github.com/amardegan/rinha-de-backend-2024-q1/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ListClientes() ([]*entity.Cliente, error) {
	clientes, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(clientes) == 0 {
		return nil, entity.ErrNaoEncontrado
	}
	return clientes, nil
}
