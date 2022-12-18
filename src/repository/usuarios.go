package repository

import (
	"database/sql"
	"fmt"
	"produto/src/models"
)

// Usuarios representa um repositorio de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, senha) values(?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos os usuarios que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nome string) ([]models.Usuario, error) {
	nome = fmt.Sprintf("%%%s%%", nome) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, senha from usuarios where nome LIKE ?",
		nome,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Senha,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuario do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, senha from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Senha,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera os dados do usuario do banco
func (repositorio Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um usuario do banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorNome busca um usuário por nome e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorNome(nome string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where nome = ?", nome)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}
