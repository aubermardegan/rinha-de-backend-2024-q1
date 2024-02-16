package main

import (
	"fmt"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
)

func main() {

	bufferCliente := entity.InitBufferClientes()

	//TODO: buscar do BD e remover esse bloco
	c1, err := entity.NewCliente(1, 100000, 0)
	if err == nil {
		bufferCliente.AddCliente(c1)
	}
	c2, err := entity.NewCliente(2, 80000, 0)
	if err == nil {
		bufferCliente.AddCliente(c2)
	}
	c3, err := entity.NewCliente(3, 1000000, 0)
	if err == nil {
		bufferCliente.AddCliente(c3)
	}
	c4, err := entity.NewCliente(4, 10000000, 0)
	if err == nil {
		bufferCliente.AddCliente(c4)
	}
	c5, err := entity.NewCliente(5, 500000, 0)
	if err == nil {
		bufferCliente.AddCliente(c5)
	}

	mc1, ok := bufferCliente.GetCliente(1)
	if ok {
		fmt.Println(mc1.Limite)
	}
}
