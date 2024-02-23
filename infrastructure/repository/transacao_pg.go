package repository

import (
	"database/sql"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
)

type TransacaoPg struct {
	db *sql.DB
}

func NewTransacaoRepository(db *sql.DB) *TransacaoPg {
	return &TransacaoPg{
		db: db,
	}
}

func (r *TransacaoPg) BeginTran() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *TransacaoPg) RollbackTran(tx *sql.Tx) error {
	return tx.Rollback()
}

func (r *TransacaoPg) CommitTran(tx *sql.Tx) error {
	return tx.Commit()
}

func (r *TransacaoPg) GetByClienteAfterId(tx *sql.Tx, clienteId, transacaoId int) ([]*entity.Transacao, error) {
	var transacoes []*entity.Transacao

	rows, err := tx.Query(`
	SELECT 
		id, valor, tipo, descricao, realizadaEm 	
	FROM transacao 
	WHERE clienteId = $1 
	  AND id > $2
	ORDER BY id
	FOR UPDATE`,
		clienteId, transacaoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transacao entity.Transacao
		err := rows.Scan(&transacao.Id, &transacao.Valor, &transacao.Tipo, &transacao.Descricao, &transacao.RealizadaEm)
		if err != nil {
			return nil, err
		}

		transacoes = append(transacoes, &transacao)
	}

	return transacoes, nil
}

func (r *TransacaoPg) GetLatestByCliente(clienteId, quantidade int) ([]*entity.Transacao, error) {
	var transacoes []*entity.Transacao

	rows, err := r.db.Query(`
	SELECT
		id, valor, tipo, descricao, realizadaEm 
	FROM transacao 
	WHERE clienteId = $1
	ORDER BY realizadaEm DESC, id DESC
	LIMIT $2`,
		clienteId, quantidade)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transacao entity.Transacao
		err := rows.Scan(&transacao.Id, &transacao.Valor, &transacao.Tipo, &transacao.Descricao, &transacao.RealizadaEm)
		if err != nil {
			return nil, err
		}

		transacoes = append(transacoes, &transacao)
	}

	return transacoes, nil
}

func (r *TransacaoPg) Create(tx *sql.Tx, clienteId int, t *entity.Transacao) (*entity.Transacao, error) {
	var transacaoId int
	err := tx.QueryRow(`
		INSERT INTO transacao (
			clienteId, valor, tipo, descricao, realizadaEm) 
		VALUES (
			$1, $2, $3, $4, $5) 
		RETURNING id`,
		clienteId, t.Valor, t.Tipo, t.Descricao, t.RealizadaEm).Scan(&transacaoId)
	if err != nil {
		return nil, err
	}

	var insertedTransacao entity.Transacao
	err = tx.QueryRow(`
	SELECT 
		id, valor, tipo, descricao, realizadaEm 
	FROM transacao 
	WHERE id = $1`,
		transacaoId).Scan(&insertedTransacao.Id, &insertedTransacao.Valor, &insertedTransacao.Tipo, &insertedTransacao.Descricao, &insertedTransacao.RealizadaEm)
	if err != nil {
		return nil, err
	}

	return &insertedTransacao, nil
}
