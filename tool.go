package login

import (
	"bytes"
	"encoding/json"
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

// encodeToUUID so it can be stored, but not in plain text.
func encodeToUUID(data string) string {
	return uuid.NewV5(uuid.NamespaceOID, data).String()
}

// hashPassword Obfuscate the password from plain text.
func hashPassword(password string) string {
	return uuid.NewV5(uuid.Nil, password).String()
}

func newJsonDecoder(data []byte) (*json.Decoder, error) {
	if len(data) == 0 || string(data) == "null" {
		return nil, fmt.Errorf(`"%s" is not a valid JSON`, string(data))
	}

	dec := json.NewDecoder(bytes.NewReader(data))

	t1, e1 := dec.Token()
	if e1 != nil {
		return nil, e1
	}
	delim, ok := t1.(json.Delim)
	if !ok || delim != '{' {
		return nil, fmt.Errorf("%v", stderr.DecodeJSONStart)
	}
	return dec, nil
}
