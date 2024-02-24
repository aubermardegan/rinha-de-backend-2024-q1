package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/amardegan/rinha-de-backend-2024-q1/api/presenter"
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
			w.WriteHeader(http.StatusNotFound)
			return
		}

		bufferClientes.RLock()
		c, ok := bufferClientes.GetCliente(intId)
		bufferClientes.RUnlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			bufferClientes.Lock()
			transacoes, err := ts.GetUltimasTransacoes(c, 10)
			bufferClientes.Unlock()
			if err != nil && !errors.Is(err, entity.ErrTransacaoNaoEncontrada) {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			toJ := &presenter.Extrato{
				Saldo: presenter.Saldo{
					Total:  c.Saldo.Valor,
					Data:   time.Now().Local().UTC(),
					Limite: c.Limite,
				},
				UltimasTransacoes: transacoes,
			}
			if err := json.NewEncoder(w).Encode(toJ); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}
