package login

import (
	"fmt"

	"github.com/google/uuid"
)

// generateId Uses UUID v7 to generate an ID.
func generateId() string {
	id, e1 := uuid.NewV7()
	if e1 != nil {
		msg := fmt.Errorf(stderr.GenUUIDv7, e1.Error())
		panic(msg)
	}
	return id.String()
}
