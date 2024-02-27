package transacao

import (
	"errors"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetUltimasTransacoes(c *entity.Cliente, quantidade int) ([]*entity.Transacao, error) {
	tx, err := s.repo.BeginTran()
	if err != nil {
		return nil, err
	}
	defer s.repo.RollbackTran(tx)

	transacoes, err := s.repo.GetByClienteAfterId(tx, c.Id, c.Saldo.UltimoIdTransacaoConferido)
	if err != nil && !errors.Is(err, entity.ErrTransacaoNaoEncontrada) {
		return nil, err
	}
	for _, transacao := range transacoes {
		atualizaSaldo(&c.Saldo, transacao)
	}

	transacoes, err = s.repo.GetLatestByCliente(c, quantidade)
	if err != nil {
		return nil, err
	}
	if len(transacoes) == 0 {
		return nil, entity.ErrTransacaoNaoEncontrada
	}

	err = s.repo.CommitTran(tx)
	if err != nil {
		return nil, err
	}

	return transacoes, nil
}

func (s *Service) CreateTransacao(c entity.Cliente, t *entity.Transacao) (int, int, error) {
	tx, err := s.repo.BeginTran()
	if err != nil {
		return 0, 0, err
	}
	defer s.repo.RollbackTran(tx)

	transacoes, err := s.repo.GetByClienteAfterId(tx, c.Id, c.Saldo.UltimoIdTransacaoConferido)
	if err != nil && !errors.Is(err, entity.ErrTransacaoNaoEncontrada) {
		return 0, 0, err
	}
	for _, transacao := range transacoes {
		atualizaSaldo(&c.Saldo, transacao)
	}

	if c.SemLimiteDisponivel(t) {
		return 0, 0, entity.ErrSemLimiteParaTransacao
	}

	insertedTransacao, err := s.repo.Create(tx, c.Id, t)
	if err != nil {
		return 0, 0, err
	}
	atualizaSaldo(&c.Saldo, insertedTransacao)

	err = s.repo.CommitTran(tx)
	if err != nil {
		return 0, 0, err
	}

	return c.Saldo.Valor, insertedTransacao.Id, nil
}

func atualizaSaldo(s *entity.Saldo, t *entity.Transacao) {
	switch t.Tipo {
	case entity.TransacaoCredito:
		s.Valor += t.Valor
	case entity.TransacaoDebito:
		s.Valor -= t.Valor
	}
	s.UltimoIdTransacaoConferido = t.Id
}
