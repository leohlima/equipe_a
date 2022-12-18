package rotas

import (
	"net/http"
	"produto/src/controllers"
)

var rotasProdutos = []Rota{
	{
		URI:                "/api/v1/product",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarProduto,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/v1/products",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarProdutos,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/v1/product/{produtoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarProduto,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/v1/product/{produtoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarProduto,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/v1/product/{produtoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarProduto,
		RequerAutenticacao: true,
	},
}
