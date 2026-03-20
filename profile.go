package login

import (
	"encoding/json"
	"fmt"

	"github.com/kohirens/storage"
)

const (
	prefixProfile = "profile/"
)

type OIDCProvider struct{}

// Profile represents a single client/user (or you), the userinfo is your PII.
// A profile can only be attached to single account. at any given time.
// There are 2 types of profiles (primary and secondary). A primary profile
// will be a profile that was made at the same time an account was made. A
// primary profile cannot move between accounts because it is set as the owner
// of the account and is bound to it. A secondary profile is a profile that was
// later added to the account, like a spouse/child/friend, for the purpose of
// sharing access to the account. A secondary profile can be removed from an account, where a
// primary profile cannot because it owns the account.
// You can think of a profile as like a profile on a game console or streaming
// service. It represents a client of that service and their access.
type Profile struct {
	ClientApp     map[string]*ClientApp    `json:"clientApp"`
	Id            string                   `json:"id"`
	Name          string                   `json:"name"`
	OIDCProviders map[string]*OIDCProvider `json:"oidcProviders"`
	UserInfo      *UserInfo                `json:"userInfo"`
}

// Save will save a profile to storage.
func (p *Profile) Save(store storage.Storage) error {
	data, e1 := json.Marshal(p)
	if e1 != nil {
		return fmt.Errorf(stderr.EncodeJSON, e1.Error())
	}

	loc := profileLocation(p.Id)

	return store.Save(loc, data)
}

// DeleteProfile will delete a profile from storage.
func DeleteProfile(id string, store storage.Storage) error {
	loc := profileLocation(id)

	if e1 := store.Remove(loc); e1 != nil {
		return e1
	}

	return nil
}

// LoadProfile will read a profile from storage.
func LoadProfile(id string, store storage.Storage) (*Profile, error) {
	loc := profileLocation(id)

	data, e1 := store.Load(loc)
	if e1 != nil {
		return nil, e1
	}

	var p *Profile

	if e2 := json.Unmarshal(data, &p); e2 != nil {
		return nil, e2
	}

	return p, nil
}

func NewProfile(name, clientAppId string, userInfo *UserInfo) *Profile {
	if name == "" {
		panic("profile name is required")
	}
	if len(name) > 100 {
		panic(fmt.Sprintf(stderr.LongProfileName, name))
	}
	validateUserInfo(userInfo)

	clientApps := make(map[string]*ClientApp, 1)
	clientApps[clientAppId] = &ClientApp{}
	return &Profile{
		ClientApp:     clientApps,
		Id:            generateId(),
		Name:          name,
		OIDCProviders: make(map[string]*OIDCProvider),
		UserInfo:      userInfo,
	}
}

func profileLocation(id string) string {
	return prefixProfile + id + filExt
}
