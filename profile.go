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

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func NewProfile(name, clientAppId string, userInfo *UserInfo) *Profile {
	if name == "" {
		panic("profile name is required")
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

func validateUserInfo(info *UserInfo) {
	// TODO: Validate email, firstname, lastname, phone with better standards.
	if info == nil {
		panic("nil UserInfo")
	}
	if info.FirstName == "" {
		panic("UserInfo FirstName is empty")
	}
	if info.LastName == "" {
		panic("UserInfo LastName is empty")
	}
	if info.Phone == "" {
		panic("UserInfo Phone is empty")
	}
	if info.Email == "" {
		panic("UserInfo Email is empty")
	}
}
