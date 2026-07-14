package login

import (
	"github.com/kohirens/stdlib/logger"
)

const (
	fileExt           = ".json"
	prefixAccountLink = "account-link/"
)

var (
	Log = &logger.Standard{}
)

func accountLinkLocation(email string) string {
	return prefixAccountLink + encodeToUUID(email) + fileExt
}
