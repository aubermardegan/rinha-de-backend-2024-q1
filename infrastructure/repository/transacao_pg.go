package repository

import (
	"context"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetLatestByCliente(ctx context.Context, db *pgxpool.Pool, c *entity.Cliente, quantidade int) ([]*entity.Transacao, error) {
	var transacoes []*entity.Transacao

	rows, err := db.Query(ctx, `
	SELECT
		id, valor, tipo, descricao, realizadaEm 
	FROM transacao 
	WHERE clienteId = $1
	ORDER BY realizadaEm DESC, id DESC
	LIMIT $2`,
		c.Id, quantidade)
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

func CreateTransacao(ctx context.Context, db *pgxpool.Pool, clienteId int, t *entity.Transacao) (int, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var saldo int
	var limite int
	err = tx.QueryRow(ctx, "SELECT saldo, limite FROM cliente WHERE id = $1 FOR UPDATE;", clienteId).Scan(&saldo, &limite)
	if err != nil {
		return 0, err
	}

	saldo = atualizaSaldo(saldo, t)
	if saldo < (limite * -1) {
		return 0, entity.ErrSemLimiteParaTransacao
	}

	batch := &pgx.Batch{}
	batch.Queue(`
	INSERT INTO transacao (
		clienteId, valor, tipo, descricao, realizadaEm) 
	VALUES (
		$1, $2, $3, $4, $5)`,
		clienteId, t.Valor, t.Tipo, t.Descricao, t.RealizadaEm)

	batch.Queue("UPDATE cliente SET saldo = $1 WHERE id = $2;", saldo, clienteId)

	br := tx.SendBatch(ctx, batch)
	_, err = br.Exec()
	if err != nil {
		return 0, err
	}
	err = br.Close()
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return saldo, nil
}

func atualizaSaldo(saldo int, t *entity.Transacao) int {
	switch t.Tipo {
	case entity.TransacaoCredito:
		saldo += t.Valor
	case entity.TransacaoDebito:
		saldo -= t.Valor
	}
	return saldo
}
