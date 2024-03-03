package transacao

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

func (s *Service) GetUltimasTransacoes(ctx context.Context, c *entity.Cliente, quantidade int) ([]*entity.Transacao, error) {

	transacoes, err := repository.GetLatestByCliente(ctx, s.repo.Pool, c, quantidade)
	if err != nil {
		return nil, err
	}
	if len(transacoes) == 0 {
		return nil, entity.ErrTransacaoNaoEncontrada
	}

	return transacoes, nil
}

func (s *Service) CreateTransacao(ctx context.Context, clienteId int, t *entity.Transacao) (int, error) {
	return repository.CreateTransacao(ctx, s.repo.Pool, clienteId, t)
}
