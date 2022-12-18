package router

import (
	"produto/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar(r *mux.Router) *mux.Router {
	return rotas.Configurar(r)
}
