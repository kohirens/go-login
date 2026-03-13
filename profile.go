package login

import (
	"encoding/json"
	"fmt"

	"github.com/kohirens/storage"
)

const (
	prefixProfile = "profile/"
)

type ClientApp struct{}
type OIDCProvider struct{}

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
