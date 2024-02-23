package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/amardegan/rinha-de-backend-2024-q1/usecase/transacao"
)

func CreateTransacao(bufferClientes *entity.BufferClientes, ts transacao.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO: Implementar
		w.WriteHeader(http.StatusNotImplemented)
	})
}

func Extrato(bufferClientes *entity.BufferClientes, ts transacao.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		bufferClientes.RLock()
		c, ok := bufferClientes.GetCliente(intId)
		bufferClientes.RUnlock()
		if ok {
			transacoes, err := ts.GetUltimasTransacoes(c.Id, 10)
			if err != nil {
				fmt.Print(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
			//TODO: criar presenter, ajustar o retorno e validar as demais regras
			if err := json.NewEncoder(w).Encode(transacoes); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}
