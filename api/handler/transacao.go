package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/amardegan/rinha-de-backend-2024-q1/api/presenter"
	"github.com/amardegan/rinha-de-backend-2024-q1/entity"
	"github.com/amardegan/rinha-de-backend-2024-q1/usecase/cliente"
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
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		c, ok := bufferClientes.GetCliente(intId)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			saldo, err := ts.CreateTransacao(c.Id, t)
			if err != nil {
				if errors.Is(err, entity.ErrSemLimiteParaTransacao) {
					w.WriteHeader(http.StatusUnprocessableEntity)
					return
				}
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			var output struct {
				Limite int `json:"limite"`
				Saldo  int `json:"saldo"`
			}
			output.Limite = c.Limite
			output.Saldo = saldo
			if err := json.NewEncoder(w).Encode(output); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}

func Extrato(bufferClientes *entity.BufferClientes, cs cliente.UseCase, ts transacao.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		c, ok := bufferClientes.GetCliente(intId)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		} else {

			cliente, err := cs.GetClienteById(c.Id)
			if err != nil && !errors.Is(err, entity.ErrClienteNaoEncontrado) {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			transacoes, err := ts.GetUltimasTransacoes(c, 10)
			if err != nil && !errors.Is(err, entity.ErrTransacaoNaoEncontrada) {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			toJ := &presenter.Extrato{
				Saldo: presenter.Saldo{
					Total:  cliente.Saldo,
					Data:   time.Now().Local().UTC(),
					Limite: cliente.Limite,
				},
				UltimasTransacoes: transacoes,
			}
			if err := json.NewEncoder(w).Encode(toJ); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}
