package transacao

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

func (s *Service) GetUltimasTransacoes(c *entity.Cliente, quantidade int) ([]*entity.Transacao, error) {

	transacoes, err := repository.GetLatestByCliente(s.repo, c, quantidade)
	if err != nil {
		return nil, err
	}
	if len(transacoes) == 0 {
		return nil, entity.ErrTransacaoNaoEncontrada
	}

	return transacoes, nil
}

func (s *Service) CreateTransacao(clienteId int, t *entity.Transacao) (int, error) {
	return repository.CreateTransacao(s.repo, clienteId, t)
}
