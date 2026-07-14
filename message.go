package login

var stderr = struct {
	AccountNotFound,
	ClientAppID,
	DecodeJSON,
	DecodeJSONKey,
	DecodeJSONStart,
	DeleteClientApp,
	EncodeJSON,
	GenUUIDv4,
	FindClientApp,
	LoadClientApp,
	LongProfileName,
	NoInfo,
	NoLogin,
	ProfileName,
	RemoveAccount,
	RemoveClientApp,
	SaveLogin,
	UserEmail,
	UserFirstName,
	UserLastName,
	UserPhone string
}{
	AccountNotFound: "account %v was not found",
	ClientAppID:     "client app ID is empty",
	DecodeJSON:      "cannot decode JSON, %v",
	DecodeJSONKey:   "expected string key in JSON object",
	DecodeJSONStart: "expected JSON object starting with '{'",
	DeleteClientApp: "cannot delete client app %v from profile %v",
	EncodeJSON:      "cannot encode JSON, %v",
	GenUUIDv4:       "cannot generate v4 uuid %v",
	FindClientApp:   "cannot find client app ID %v listed in the profile %v",
	LoadClientApp:   "cannot load client app ID %v from storage %v",
	LongProfileName: "profile name %v was too long, profile names are limited to 100 characters",
	NoInfo:          "nil UserInfo",
	NoLogin:         "login info %v was not found",
	ProfileName:     "profile name is required",
	RemoveAccount:   "cannot remove account %v",
	RemoveClientApp: "cannot find the client app ID %v in the profile %v in order to remove",
	UserEmail:       "UserInfo Email is empty",
	UserFirstName:   "UserInfo FirstName is empty",
	UserLastName:    "UserInfo LastName is empty",
	UserPhone:       "UserInfo Phone is empty",
}

var stdout = struct {
	ClientAppLoad string
}{
	ClientAppLoad: "Loading client app %v",
}
