package login

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kohirens/storage"
)

const (
	prefixAccount = "account/"
)

type Account struct {
	Id       string                 `json:"id"`
	Owner    string                 `json:"owner"`
	Profiles map[string]*SubProfile `json:"profiles"`
}

// Save to storage, a.k.a. Creat/Update.
func (act *Account) Save(store storage.Storage) error {
	data, e1 := json.Marshal(act)
	if e1 != nil {
		return e1
	}

	loc := accountLocation(act.Id)

	return store.Save(loc, data)
}

type SubProfile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// DeleteAccount from storage a.k.a the D in CRUD.
func DeleteAccount(id string, store storage.Storage) error {
	loc := accountLocation(id)

	if e := store.Remove(loc); e != nil {
		return fmt.Errorf(stderr.RemoveAccount, e.Error())
	}

	return nil
}

// AccountLink link an account by email and password.
type AccountLink struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	AccountID string `json:"accountID"`
}

// FindAccount will read a login from storage.
func FindAccount(email, password string, store storage.Storage) (*Account, error) {
	loc := accountLinkLocation(email)

	data, e1 := store.Load(loc)
	if e1 != nil {
		return nil, e1
	}

	var l *AccountLink

	if e2 := json.Unmarshal(data, &l); e2 != nil {
		return nil, e2
	}

	// Verify the correct password was entered.
	if l.Password != hashPassword(password) {
		return nil, errors.New("invalid password")
	}

	return LoadAccount(l.AccountID, store)
}

// LoadAccount from storage a.k.a READ.
func LoadAccount(id string, store storage.Storage) (*Account, error) {
	loc := accountLocation(id)

	data, e1 := store.Load(loc)
	if e1 != nil {
		return nil, fmt.Errorf(stderr.AccountNotFound, e1.Error())
	}

	var act *Account

	if e2 := json.Unmarshal(data, &act); e2 != nil {
		return nil, fmt.Errorf(stderr.DecodeJSON, e2.Error())
	}

	return act, nil
}

func NewAccount(profileId, profileName string) *Account {
	return &Account{
		Id: generateId(),
		Profiles: map[string]*SubProfile{
			profileId: {
				Id:   profileId,
				Name: profileName,
			},
		},
		Owner: profileId,
	}
}

func accountLocation(id string) string {
	return prefixAccount + id + fileExt
}
