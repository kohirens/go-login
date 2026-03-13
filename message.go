package login

var stderr = struct {
	AccountNotFound,
	DecodeJSON,
	EncodeJSON,
	GenUUIDv7,
	LongProfileName,
	NoInfo,
	RemoveAccount,
	SaveLogin,
	UserEmail,
	UserFirstName,
	UserLastName,
	UserPhone string
}{
	AccountNotFound: "account %v was not found",
	DecodeJSON:      "cannot decode JSON, %v",
	EncodeJSON:      "cannot encode JSON, %v",
	GenUUIDv7:       "cannot generate v7 uuid %v",
	LongProfileName: "profile name %v was too long, profile names are limited to 100 characters",
	NoInfo:          "nil UserInfo",
	RemoveAccount:   "cannot remove account %v",
	UserEmail:       "UserInfo Email is empty",
	UserFirstName:   "UserInfo FirstName is empty",
	UserLastName:    "UserInfo LastName is empty",
	UserPhone:       "UserInfo Phone is empty",
}
