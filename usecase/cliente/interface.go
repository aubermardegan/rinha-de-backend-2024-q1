package cliente

import "github.com/amardegan/rinha-de-backend-2024-q1/entity"

type Reader interface {
	List() ([]*entity.Cliente, error)
	Get() (*entity.Cliente, error)
}

type Repository interface {
	Reader
}

type UseCase interface {
	ListClientes() ([]*entity.Cliente, error)
	GetClienteById(int) (*entity.Cliente, error)
}
