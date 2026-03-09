package login

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kohirens/storage"
)

const (
	filExt      = ".json"
	prefixLogin = "login/"
)

type Login struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	LastDate time.Time `json:"lastDate"`
}

// Save to storage.
func (l *Login) Save(store storage.Storage) error {
	data, e1 := json.Marshal(l)
	if e1 != nil {
		return e1
	}

	loc := loginLocation(l.Email)

	return store.Save(loc, data)
}

// LoadLogin from storage.
func LoadLogin(email, password string, store storage.Storage) (*Login, error) {
	loc := loginLocation(email)

	data, e1 := store.Load(loc)
	if e1 != nil {
		return nil, e1
	}

	var l *Login

	if e2 := json.Unmarshal(data, &l); e2 != nil {
		return nil, e2
	}

	// Verify the correct password was entered.
	if l.Password != hashPassword(password) {
		return nil, errors.New("invalid password")
	}

	return l, nil
}

func NewLogin(email string, password string) *Login {
	return &Login{
		Email:    email,
		Password: hashPassword(password),
		LastDate: time.Now().UTC(),
	}
}

func hashPassword(password string) string {
	return uuid.NewSHA1(uuid.Nil, []byte(password)).String()
}

func loginLocation(email string) string {
	hash := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(email))
	return prefixLogin + hash.String() + filExt
}
