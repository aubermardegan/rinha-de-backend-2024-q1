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

	bufferCliente := entity.InitBufferClientes()

	ctx := context.Background()
	db, err := repository.InitPostgreSQL(ctx)
	if err != nil {
		panic(err)
	}

	cs := cliente.NewService(db)
	ts := transacao.NewService(db)

	clientes, err := cs.ListClientes(ctx)
	if err != nil {
		panic(err)
	}

	for _, c := range clientes {
		bufferCliente.AddCliente(c)
	}

	api.InitAPI(ctx, &bufferCliente, cs, ts)
}
