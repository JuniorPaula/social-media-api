package models

import (
	"errors"
	"strings"
	"time"
)

type Users struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *Users) HandlerUser() error {
	if err := u.validator(); err != nil {
		return err
	}

	u.formatFields()
	return nil
}

func (u *Users) validator() error {
	if u.Name == "" {
		return errors.New("o campo nome é obrigatório")
	}
	if u.Nickname == "" {
		return errors.New("o campo nickname é obrigatório")
	}
	if u.Email == "" {
		return errors.New("o campo email é obrigatório")
	}
	return nil
}

func (u *Users) formatFields() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)
}
