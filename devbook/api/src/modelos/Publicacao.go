package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar a publicação recebida
func (Publicacao *Publicacao) Preparar() error {
	if erro := Publicacao.validar(); erro != nil {
		return erro
	}
	Publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O titulo é obrigatorio e nao pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("O conteudo é obrigatorio e nao pode estar em branco")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
