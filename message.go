package login

var stderr = struct {
	AccountNotFound,
	DecodeJSON,
	EncodeJSON,
	GenUUIDv7,
	RemoveAccount string
}{
	AccountNotFound: "account %v was not found",
	DecodeJSON:      "cannot decode JSON, %v",
	EncodeJSON:      "cannot encode JSON, %v",
	GenUUIDv7:       "cannot generate v7 uuid %v",
	RemoveAccount:   "cannot remove account %v",
}
