package login

var stderr = struct {
	AccountNotFound,
	DecodeJSON,
	DeleteClientApp,
	EncodeJSON,
	GenUUIDv7,
	FindClientApp,
	LongProfileName,
	NoInfo,
	RemoveAccount,
	RemoveClientApp,
	SaveLogin,
	UserEmail,
	UserFirstName,
	UserLastName,
	UserPhone string
}{
	AccountNotFound: "account %v was not found",
	DecodeJSON:      "cannot decode JSON, %v",
	DeleteClientApp: "cannot delete client app %v from profile %v",
	EncodeJSON:      "cannot encode JSON, %v",
	GenUUIDv7:       "cannot generate v7 uuid %v",
	FindClientApp:   "cannot find client app ID %v listed in the profile %v",
	LongProfileName: "profile name %v was too long, profile names are limited to 100 characters",
	NoInfo:          "nil UserInfo",
	RemoveAccount:   "cannot remove account %v",
	RemoveClientApp: "cannot find the client app ID %v in the profile %v in order to remove",
	UserEmail:       "UserInfo Email is empty",
	UserFirstName:   "UserInfo FirstName is empty",
	UserLastName:    "UserInfo LastName is empty",
	UserPhone:       "UserInfo Phone is empty",
}
