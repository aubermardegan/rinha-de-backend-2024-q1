package repository

import (
	"database/sql"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
)

func ListClientes(db *sql.DB) ([]*entity.Cliente, error) {
	var clientes []*entity.Cliente

	rows, err := db.Query("SELECT id, limite, saldo FROM cliente")
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

func GetClienteById(db *sql.DB, clienteId int) (*entity.Cliente, error) {
	var cliente entity.Cliente

	err := db.QueryRow("SELECT id, limite, saldo FROM cliente WHERE Id = $1", clienteId).Scan(&cliente.Id, &cliente.Limite, &cliente.Saldo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrClienteNaoEncontrado
		}
		return nil, err
	}

	return &cliente, nil
}
