package models

import (
	"erickramos-go/src/security"
	"errors"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID                    uint64 `json:"id,omitempty"`
	FullName              string `json:"fullName,omitempty"`
	Username              string `json:"username,omitempty"`
	Password              string `json:"password,omitempty"`
	AccountNonExpired     bool   `json:"accountNonExpired,omitempty"`
	AccountNonLocked      bool   `json:"accountNonLocked,omitempty"`
	CredentialsNonExpired bool   `json:"credentialsNonExpired,omitempty"`
	Enabled               bool   `json:"enabled,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	if err := usuario.formatar(etapa); err != nil {
		return err
	}

	return nil
}

func (usuario Usuario) validar(etapa string) error {
	if usuario.FullName == "" {
		return errors.New("O FullName é obrigatório e não pode estar em branco")
	}

	if usuario.Password == "" {
		return errors.New("O Password é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {

	if etapa == "cadastro" {
		senhaComHash, err := security.Hash(usuario.Password)
		if err != nil {
			return err
		}

		usuario.Password = string(senhaComHash)
	}

	return nil
}
