package main

import (
	"fmt"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/amardegan/rinha-de-backend-2024-q1/infrastructure/repository"
	"github.com/amardegan/rinha-de-backend-2024-q1/usecase/cliente"
)

func main() {

	bufferCliente := entity.InitBufferClientes()

	db, err := repository.InitPostgreSQL()
	if err != nil {
		panic(err)
	}

	cs := cliente.NewService(repository.NewClienteRepository(db))

	clientes, err := cs.ListClientes()
	if err != nil {
		panic(err)
	}

	for _, c := range clientes {
		bufferCliente.AddCliente(c)
	}

	mc1, ok := bufferCliente.GetCliente(1)
	if ok {
		fmt.Println(mc1.Limite)
	}
}
