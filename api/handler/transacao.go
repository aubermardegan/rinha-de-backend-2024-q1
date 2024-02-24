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

		id := r.PathValue("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var input struct {
			Valor     int    `json:"valor"`
			Tipo      string `json:"tipo"`
			Descricao string `json:"descricao"`
		}
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t, err := entity.NewTransacao(input.Valor, input.Tipo, input.Descricao)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		bufferClientes.RLock()
		c, ok := bufferClientes.GetCliente(intId)
		bufferClientes.RUnlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			bufferClientes.Lock()
			err := ts.CreateTransacao(c, t)
			bufferClientes.Unlock()
			if err != nil {
				if errors.Is(err, entity.ErrSemLimiteParaTransacao) {
					w.WriteHeader(http.StatusUnprocessableEntity)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			var output struct {
				Limite int `json:"limite"`
				Saldo  int `json:"saldo"`
			}
			output.Limite = c.Limite
			output.Saldo = c.Saldo.Valor
			if err := json.NewEncoder(w).Encode(output); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
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
