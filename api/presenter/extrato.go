package presenter

import (
	"time"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
)

type Extrato struct {
	Saldo             Saldo               `json:"saldo"`
	UltimasTransacoes []*entity.Transacao `json:"ultimas_transacoes"`
}

type Saldo struct {
	Total  int       `json:"total"`
	Data   time.Time `json:"data_extrato"`
	Limite int       `json:"limite"`
}
