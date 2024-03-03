package repository

import (
	"context"
	"database/sql"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ListClientes(ctx context.Context, db *pgxpool.Pool) ([]*entity.Cliente, error) {
	var clientes []*entity.Cliente

	rows, err := db.Query(ctx, "SELECT id, limite, saldo FROM cliente")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cliente entity.Cliente
		err := rows.Scan(&cliente.Id, &cliente.Limite, &cliente.Saldo)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, &cliente)
	}

	return clientes, nil
}

func GetClienteById(ctx context.Context, db *pgxpool.Pool, clienteId int) (*entity.Cliente, error) {
	var cliente *entity.Cliente
	rows, err := db.Query(ctx, "SELECT id, limite, saldo FROM cliente WHERE Id = $1", clienteId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cliente, err = pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[entity.Cliente])
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrClienteNaoEncontrado
		}
		return nil, err
	}

	return cliente, nil
}
