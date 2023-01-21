package models

import (
	"api/src/auth"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Users struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *Users) HandlerUser(step string) error {
	if err := u.validator(step); err != nil {
		return err
	}

	if err := u.formatFields(step); err != nil {
		return err
	}

	return nil
}

func (u *Users) validator(step string) error {
	if u.Name == "" {
		return errors.New("o campo nome é obrigatório")
	}
	if u.Nickname == "" {
		return errors.New("o campo nickname é obrigatório")
	}
	if u.Email == "" {
		return errors.New("o campo email é obrigatório")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("e-mail inválido")
	}

	if step == "create" && u.Password == "" {
		return errors.New("o campo senha é obrigatório")
	}
	return nil
}

func (u *Users) formatFields(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)

	if step == "create" {
		passwordHashed, err := auth.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(passwordHashed)
	}
	return nil
}
