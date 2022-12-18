package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"produto/response"
	"produto/src/db"
	"produto/src/models"
	"produto/src/repository"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarProduto configura a rota para criar um produto no db de dados
func CriarProduto(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var produtos models.Produto

	if erro = json.Unmarshal(corpoRequest, &produtos); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = produtos.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeProdutos(db)
	produtos.ID, erro = repositorio.Criar(produtos)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, produtos)
}

// BuscarProdutos configura a rota para buscar produtos no db de dados
func BuscarProdutos(w http.ResponseWriter, r *http.Request) {
	nomeOuCodigo := strings.ToLower(r.URL.Query().Get("products"))

	db, erro := db.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeProdutos(db)
	produtos, erro := repositorio.Buscar(nomeOuCodigo)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, produtos)
}

// BuscarProduto configura a rota para criar um produto no db de dados
func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	produtoID, erro := strconv.ParseUint(parametros["produtoId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeProdutos(db)
	usuario, erro := repositorio.BuscarPorID(produtoID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, usuario)
}

// Atualizar configura a rota para atualizar um produto no db de dados
func AtualizarProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	produtoID, erro := strconv.ParseUint(parametros["produtoId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var produto models.Produto
	if erro = json.Unmarshal(corpoRequisicao, &produto); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = produto.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeProdutos(db)
	if erro = repositorio.Atualizar(produtoID, produto); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeletarProduto configura a rota para deletar um produto no db de dados
func DeletarProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	produtoID, erro := strconv.ParseUint(parametros["produtoId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeProdutos(db)
	if erro = repositorio.Deletar(produtoID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
