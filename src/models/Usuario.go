package models

import (
	"errors"
	"produto/src/security"
	"strings"
)

// Usuario representa uma entidada de usuario
type Usuario struct {
	ID    uint64 `json:"id,omitempty"`
	Nome  string `json:"username,omitempty"`
	Senha string `json:"password,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)

	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
