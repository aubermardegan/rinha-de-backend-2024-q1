package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/amardegan/rinha-de-backend-2024-q1/api/handler"
	"github.com/amardegan/rinha-de-backend-2024-q1/config"
	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/amardegan/rinha-de-backend-2024-q1/usecase/cliente"
	"github.com/amardegan/rinha-de-backend-2024-q1/usecase/transacao"
)

func InitAPI(ctx context.Context, bufferClientes *entity.BufferClientes, cs cliente.UseCase, ts transacao.UseCase) {
	mux := http.NewServeMux()

	mux.Handle("GET /clientes/{id}/extrato", handler.Extrato(ctx, bufferClientes, cs, ts))
	mux.Handle("POST /clientes/{id}/transacoes", handler.CreateTransacao(ctx, bufferClientes, ts))

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), mux)
	if err != nil {
		panic(err)
	}
}
