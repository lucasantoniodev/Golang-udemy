package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota Estrutura de todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar colcoa todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {

	rotas := rotasUsuarios // Slice de rotas (Apenas guardando em uma variável, poderia passar direto para a iteração...)
	rotas = append(rotas, rotaLogin)
	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo) // Agora o *mux.Router tem uma HandleFunc registrada
		}
	}

	return r // Retornando a *mux.Router com as rotas configuradas
}
