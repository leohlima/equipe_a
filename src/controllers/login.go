package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"produto/response"
	"produto/src/authentication"
	"produto/src/db"
	"produto/src/models"
	"produto/src/repository"
	"produto/src/security"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorNome(usuario.Nome)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	Result_token := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	err := json.NewEncoder(w).Encode(Result_token)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
}
