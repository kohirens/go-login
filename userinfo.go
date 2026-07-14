package login

import (
	"regexp"

	"github.com/kohirens/sso/oidc"
)

// UserInfo is the personally identifiable information of a client/users
// profile. It MUST be kept secure at all times when it is not required for
// processing.
type UserInfo struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Locale    string `json:"locale"`
}

func NewUserInfo(email, firstName, lastName, phone string) *UserInfo {
	u := &UserInfo{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		ID:        generateId(),
	}

	validateUserInfo(u)

	return u
}

func NewUserByProvider(ui oidc.UserInfo) *UserInfo {
	u := &UserInfo{
		Email:     ui.Email(),
		FirstName: ui.FirstName(),
		LastName:  ui.LastName(),
		ID:        generateId(),
	}

	if p := ui.Phone(); p != "" {
		u.Phone = p
	}

	validateUserInfo(u)

	return u
}

func validateUserInfo(info *UserInfo) {
	// TODO: Validate email, firstname, lastname, phone with better standards.
	if info == nil {
		panic(stderr.NoInfo)
	}
	if info.FirstName == "" && len(info.FirstName) < 100 {
		panic(stderr.UserFirstName)
	}
	if info.LastName == "" && len(info.FirstName) < 100 {
		panic(stderr.UserLastName)
	}
	// TODO: Validate phone number based on locale
	if info.Phone != "" {
		if info.Locale == "en-US" {
			re := regexp.MustCompile("^[0-9]{3}-[0-9]{3}-[0-9]{4}$")
			if !re.MatchString(info.Phone) {
				panic(stderr.UserPhone)
			}
		}

	}

	if info.Email == "" {
		panic(stderr.UserEmail)
	}
}
