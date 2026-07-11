package login

import (
	"fmt"

	"github.com/gofrs/uuid/v5"
)

// generateId Uses UUID v7 to generate an ID.
func generateId() string {
	id, e1 := uuid.NewV4()
	if e1 != nil {
		msg := fmt.Errorf(stderr.GenUUIDv4, e1.Error())
		panic(msg)
	}
	return id.String()
}
