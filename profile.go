package login

import (
	"encoding/json"

	"github.com/kohirens/storage"
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
		return e1
	}

	return store.Save(p.Id+".json", data)
}

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func NewProfile(name, clientAppId string, userInfo *UserInfo) *Profile {
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

func validateUserInfo(info *UserInfo) {}
