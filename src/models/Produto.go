package models

import (
	"errors"
	"strings"
	"time"
)

// Produto representa uma entidada de produto
type Produto struct {
	ID     uint64    `json:"id,omitempty"`
	Nome   string    `json:"name,omitempty"`
	Codigo string    `json:"code,omitempty"`
	Valor  float64   `json:"price,omitempty"`
	DataC  time.Time `json:"-"`
	DataU  time.Time `json:"-"`
}

type ListaProdutos struct {
	Lista []Produto `json:"list,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o produto recebido
func (produto *Produto) Preparar(etapa string) error {
	if erro := produto.validar(etapa); erro != nil {
		return erro
	}
	produto.formatar()
	return nil
}

func (produto Produto) validar(etapa string) error {
	if produto.Nome == "" {
		return errors.New("o campo nome é obrigatório")
	}
	if produto.Codigo == "" {
		return errors.New("o campo código é obrigatório e não pode estar em branco ou ser zero")
	}
	if produto.Valor == 0 {
		return errors.New("o valor é obrigatório e não e não pode estar em branco ou ser zero")
	}

	return nil
}

func (produto *Produto) formatar() {
	produto.Nome = strings.TrimSpace(produto.Nome)
}
