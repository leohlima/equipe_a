package rotas

import (
	"net/http"
	"produto/src/controllers"
)

var rotaLogin = Rota{
	URI:                "/api/v1/user/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
