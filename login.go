package login

import (
	"github.com/kohirens/stdlib/logger"
)

const (
	fileExt           = ".json"
	prefixLogin       = "login/"
	prefixAccountLink = "account-link/"
	prefixClientApp   = "client-app/"
)

var (
	Log = &logger.Standard{}
)

func accountLinkLocation(email string) string {
	return prefixAccountLink + encodeToUUID(email) + fileExt
}
