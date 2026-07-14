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
	Owner    string                 `json:"owner"`
	Profiles map[string]*SubProfile `json:"profiles"`
	id       string
}

func (act *Account) ID() string {
	return act.id
}

func (act *Account) Save(store storage.Storage) error {
	data, e1 := json.Marshal(act)
	if e1 != nil {
		return e1
	}

	loc := accountLocation(act.id)

	return store.Save(loc, data)
}

func (act *Account) MarshalJSON() ([]byte, error) {
	retVal := []byte(act.String())
	return retVal, nil
}

func (act *Account) String() string {
	jsonString := `"id":"` + act.id + `",`
	jsonString += `"owner":"` + act.Owner + `",`
	profiles, e1 := json.Marshal(act.Profiles)
	if e1 != nil {
		panic(e1)
	}
	jsonString += `"profiles":` + string(profiles) + ``
	return "{" + jsonString + "}"
}

func (act *Account) UnmarshalJSON(data []byte) error {
	mar := func(d *json.Decoder, key string, obj any) error {
		switch key {
		case "id":
			if e := d.Decode(&act.id); e != nil {
				return fmt.Errorf(stderr.DecodeJSONField, key, e)
			}
		case "profiles":
			if e := d.Decode(&act.Profiles); e != nil {
				return fmt.Errorf("cannot decode %v: %w", key, e)
			}
		case "owner":
			if e := d.Decode(&act.Owner); e != nil {
				return fmt.Errorf("cannot decode %v: %w", key, e)
			}
		default:
			var discard json.RawMessage
			if e := d.Decode(&discard); e != nil {
				return fmt.Errorf("failed to skip unknown field %s: %w", key, e)
			}
		}
		return nil
	}

	return unmarshal(data, act, mar)
}

// DeleteAccount from storage a.k.a the D in CRUD.
func DeleteAccount(id string, store storage.Storage) error {
	loc := accountLocation(id)

	if e := store.Remove(loc); e != nil {
		return fmt.Errorf(stderr.RemoveAccount, e.Error())
	}

	return nil
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
		return nil, errors.New(stderr.Password)
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

	act := &Account{}

	if e2 := act.UnmarshalJSON(data); e2 != nil {
		return nil, fmt.Errorf(stderr.DecodeJSON, e2.Error())
	}

	return act, nil
}

func NewAccount(profileId, profileName string) *Account {
	return &Account{
		id: generateId(),
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
