package repository

import (
	"database/sql"
	"fmt"
	"produto/src/models"
)

// Produtos representa um repositorio de usuário
type Produtos struct {
	db *sql.DB
}

// NovoRepositoDeProdutos cria um repositorio de usuários
func NovoRepositorioDeProdutos(db *sql.DB) *Produtos {
	return &Produtos{db}
}

// Criar insere um produto no banco de dados
func (repositorio Produtos) Criar(produto models.Produto) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into produtos (nome, codigo, valor) values(?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(produto.Nome, produto.Codigo, produto.Valor)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os produtos que atendem um filtro de nome ou codigo
func (repositorio Produtos) Buscar(nomeOuCodigo string) (lista_produtos models.ListaProdutos, err error) {
	nomeOuCodigo = fmt.Sprintf("%%%s%%", nomeOuCodigo)

	linhas, erro := repositorio.db.Query(
		"select id, nome, codigo, valor from produtos where nome LIKE ? or codigo LIKE ?",
		nomeOuCodigo, nomeOuCodigo,
	)
	if erro != nil {
		return lista_produtos, err
	}
	defer linhas.Close()

	for linhas.Next() {
		var produto models.Produto

		if erro = linhas.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Codigo,
			&produto.Valor,
		); erro != nil {
			return lista_produtos, err
		}

		lista_produtos.Lista = append(lista_produtos.Lista, produto)
	}

	return lista_produtos, nil
}

// BuscarPorID traz um usuario do banco de dados
func (repository Produtos) BuscarPorID(ID uint64) (models.Produto, error) {
	linhas, erro := repository.db.Query(
		"select id, nome, codigo, valor, dataC from produtos where id = ?",
		ID,
	)
	if erro != nil {
		return models.Produto{}, erro
	}
	defer linhas.Close()

	var produto models.Produto

	if linhas.Next() {
		if erro = linhas.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Codigo,
			&produto.Valor,
			&produto.DataC,
		); erro != nil {
			return models.Produto{}, erro
		}
	}

	return produto, nil
}

// Atualizar altera os dados do produto do banco
func (repository Produtos) Atualizar(ID uint64, produto models.Produto) error {
	statement, erro := repository.db.Prepare(
		"update produtos set nome = ? , valor = ?, codigo = ?, dataU = current_timestamp() where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(produto.Nome, produto.Valor, produto.Codigo, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de u produto do banco de dados
func (repositorio Produtos) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from produtos where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
