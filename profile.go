package login

import "github.com/kohirens/storage"

type Device struct{}
type OIDCProvider struct{}

type Profile struct {
	Devices       map[string]*Device       `json:"devices"`
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

func NewProfile(deviceId string, userInfo *UserInfo) *Profile {
	validateUserInfo(userInfo)

	devices := make(map[string]*Device, 1)
	return &Profile{
		Id:            generateId(),
		UserInfo:      userInfo,
		Devices:       devices,
		OIDCProviders: make(map[string]*OIDCProvider),
	}
}

func validateUserInfo(info *UserInfo) {}
