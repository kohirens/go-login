package login

import (
	"github.com/google/uuid"
	"github.com/kohirens/storage"
)

const (
	prefixProfileMap = "access-map/"
)

// SaveProfileMap will save an access map entry to storage.
func SaveProfileMap(id, profileId string, store storage.Storage) error {
	loc := profileMapLocation(id)

	return store.Save(loc, []byte(profileId))
}

// DeleteProfileMap will remove an access map entry from storage.
func DeleteProfileMap(email string, store storage.Storage) error {
	loc := profileMapLocation(email)

	if e := store.Remove(loc); e != nil {
		return e
	}

	return nil
}

// LoadProfileMap will read a login from storage.
func LoadProfileMap(id string, store storage.Storage) (string, error) {
	loc := profileMapLocation(id)

	data, e1 := store.Load(loc)
	if e1 != nil {
		return "", e1
	}

	return string(data), nil
}

func profileMapLocation(id string) string {
	hash := uuid.NewSHA1(uuid.Nil, []byte(id))
	return prefixProfileMap + hash.String() + filExt
}
