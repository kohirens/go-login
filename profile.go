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

func (p *Profile) Save(store storage.Storage) error {

	return nil
}

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func NewProfile(clientAppId string, userInfo *UserInfo) *Profile {
	validateUserInfo(userInfo)

	clientApps := make(map[string]*ClientApp, 1)
	clientApps[clientAppId] = &ClientApp{}
	return &Profile{
		Id:            generateId(),
		UserInfo:      userInfo,
		ClientApp:     clientApps,
		OIDCProviders: make(map[string]*OIDCProvider),
	}
}

func validateUserInfo(info *UserInfo) {}
