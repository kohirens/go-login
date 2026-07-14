package login

var stderr = struct {
	AccountNotFound,
	ClientAppID,
	DecodeJSON,
	DecodeJSONField,
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
	Password,
	ProfileName,
	RemoveAccount,
	RemoveClientApp,
	UserEmail,
	UserFirstName,
	UserLastName,
	UserPhone string
}{
	AccountNotFound: "account %v was not found",
	ClientAppID:     "client app ID is empty",
	DecodeJSON:      "cannot decode JSON, %v",
	DecodeJSONField: "cannot decode %v: %w",
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
	Password:        "invalid password",
	ProfileName:     "profile name is required",
	RemoveAccount:   "cannot remove account %v",
	RemoveClientApp: "cannot find the client app ID %v in the profile %v in order to remove",
	UserEmail:       "user email is empty",
	UserFirstName:   "user first name is empty",
	UserLastName:    "user last name is empty",
	UserPhone:       "invalid user phone number format",
}

var stdout = struct {
	ClientAppLoad string
}{
	ClientAppLoad: "Loading client app %v",
}
