package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID         uint64    `json:"id,omitempty"`
	Titulo     string    `json:"titulo,omitempty"`
	Conteudo   string    `json:"conteudo,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Curtidas   uint64    `json:"curtidas"`
	CriadoEm   time.Time `json:"criado_em,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O titulo é obrigatório e não pode ser em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O conteudo é obrigatório e não pode ser em branco")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Titulo)
}
