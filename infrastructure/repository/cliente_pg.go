package repository

import (
	"database/sql"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
)

type ClientePg struct {
	db *sql.DB
}

func NewClienteRepository(db *sql.DB) *ClientePg {
	return &ClientePg{
		db: db,
	}
}

func (r *ClientePg) List() ([]*entity.Cliente, error) {
	var clientes []*entity.Cliente

	rows, err := r.db.Query("SELECT id, limite, saldoInicial FROM cliente")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cliente entity.Cliente
		var saldoInicial int
		err := rows.Scan(&cliente.Id, &cliente.Limite, &saldoInicial)
		if err != nil {
			return nil, err
		}

		cliente.Saldo.Valor = saldoInicial
		cliente.Saldo.UltimoIdTransacaoConferido = 0

		clientes = append(clientes, &cliente)
	}

	return clientes, nil
}
