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

// Save to storage.
func (p *Profile) Save(store storage.Storage) error {
	data, e1 := json.Marshal(p)
	if e1 != nil {
		return fmt.Errorf(stderr.EncodeJSON, e1.Error())
	}

	loc := profileLocation(p.Id)

	return store.Save(loc, data)
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
