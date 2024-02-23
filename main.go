package main

import (
	"context"

	"github.com/amardegan/rinha-de-backend-2024-q1/api"
	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/amardegan/rinha-de-backend-2024-q1/infrastructure/repository"
	"github.com/amardegan/rinha-de-backend-2024-q1/usecase/cliente"
	"github.com/amardegan/rinha-de-backend-2024-q1/usecase/transacao"
)

func main() {

	ctx := context.Background()
	bufferCliente := entity.InitBufferClientes()

	db, err := repository.InitPostgreSQL()
	if err != nil {
		panic(err)
	}

	cs := cliente.NewService(repository.NewClienteRepository(db))
	ts := transacao.NewService(repository.NewTransacaoRepository(db))

	clientes, err := cs.ListClientes()
	if err != nil {
		panic(err)
	}

	for _, c := range clientes {
		bufferCliente.AddCliente(c)
	}

	api.InitAPI(ctx, &bufferCliente, ts)
}
