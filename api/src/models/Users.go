package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representes a user structure
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Passwd    string    `json:"passwd,omitempty"`
	CreatedIn time.Time `json:"createdin,omitempty"`
}

// Prepare will validate and format received data
func (u *User) Prepare(step string) error {

	if err := u.validate(step); err != nil {
		return err
	}
	if err := u.format(step); err != nil {
		return err
	}
	return nil
}

func (u *User) validate(step string) error {

	if u.Name == "" {

		return errors.New("Name is required")

	}

	if u.Nick == "" {
		return errors.New("Nick is required")

	}
	if u.Email == "" {
		return errors.New("Email is required")

	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Email Address Format is Invalid")
	}

	if step == "create" && u.Passwd == "" {
		return errors.New("Password is required")

	}

	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if step == "create" {
		passwdHash, err := security.Hash(u.Passwd)
		if err != nil {
			return err
		}
		u.Passwd = string(passwdHash)
	}
	return nil
}
