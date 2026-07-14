package login

import (
	"github.com/kohirens/stdlib/logger"
)

const (
	fileExt           = ".json"
	prefixAccountLink = "account-link/"
	prefixClientApp   = "client-app/"
)

var (
	Log = &logger.Standard{}
)

func accountLinkLocation(email string) string {
	return prefixAccountLink + encodeToUUID(email) + fileExt
}

// clientAppLocation containing login information.
func clientAppLocation(id string) string {
	return prefixClientApp + id + ".json"
}
