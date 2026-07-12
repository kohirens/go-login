package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kohirens/sso/oidc"
	"github.com/kohirens/storage"
	"github.com/mileusna/useragent"
)

// ClientApp represents the client's application used to access this application
// such as a browser or other application that can view web applications.
type ClientApp struct {
	// id is a GUID generated for the browser since they do have one.
	id string
	// LastActivity indicates when the client last used this browser.
	LastActivity time.Time `json:"last_activity"`
	// Provider will be set to a provider chosen by the user to log in.
	Provider oidc.Provider `json:"provider"`
	// Meta is the user agent info of the clients' browser.
	Meta *useragent.UserAgent `json:"meta"`
}

// DeleteClientApp removes previously registered client app info from storage.
func DeleteClientApp(id string, store storage.Storage) error {
	if len(id) == 0 {
		return fmt.Errorf("%v", stderr.ClientAppID)
	}

	filename := clientAppLocation(id)

	Log.Infof(stdout.ClientAppLoad, filename)

	if e1 := store.Remove(filename); e1 != nil {
		return fmt.Errorf(stderr.LoadClientApp, e1.Error())
	}

	return nil
}

// LoadClientApp retrieve previously registered client app info from storage.
func LoadClientApp(id string, store storage.Storage) (*ClientApp, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("%v", stderr.ClientAppID)
	}

	filename := clientAppLocation(id)

	Log.Infof(stdout.ClientAppLoad, filename)

	data, e1 := store.Load(filename)
	if e1 != nil {
		return nil, fmt.Errorf(stderr.LoadClientApp, e1.Error())
	}

	var ca *ClientApp
	if e2 := json.Unmarshal(data, ca); e2 != nil {
		return nil, fmt.Errorf(stderr.DecodeJSON, e2.Error())
	}

	return ca, nil
}

// RegisterClientApp saves the new user agent to storage and return it.
//
//	NOTE: This is the only time the user agent is set in a login.
func RegisterClientApp(userAgent string, provider oidc.Provider, store storage.Storage) (*ClientApp, error) {
	ua := useragent.Parse(userAgent)
	l := &ClientApp{
		id:           encodeToUUID(userAgent),
		LastActivity: time.Now().UTC(),
		Provider:     provider,
		Meta:         &ua,
	}

	if e := save(l, store); e != nil {
		return nil, e
	}

	return l, nil
}

func (ca *ClientApp) String() string {
	jsonString := `"id": "` + ca.id + `",`
	jsonString += `"lastActivity": "` + ca.LastActivity.String() + `",`
	if ca.Provider != nil {
		jsonString += `"provider": "` + ca.Provider.String() + `",`
	}
	jsonString += `"meta": "` + ca.Meta.String + `"`
	return "{" + jsonString + "}"
}

func (ca *ClientApp) Unmarshal(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))

	t1, e1 := dec.Token()
	if e1 != nil {
		return e1
	}
	delim, ok := t1.(json.Delim)
	if !ok || delim != '{' {
		return fmt.Errorf("%v", stderr.DecodeJSONStart)
	}

	// Process all remaining tokens
	for dec.More() {
		t, e2 := dec.Token()
		if e2 != nil {
			return e2
		}

		key, k := t.(string)
		if !k {
			return fmt.Errorf("%v", stderr.DecodeJSONKey)
		}

		switch key {
		case "id":
			if e := dec.Decode(&ca.id); e != nil {
				return fmt.Errorf("cannot decode %v: %w", key, e)
			}
		case "lastActivity":
			if e := dec.Decode(&ca.LastActivity); e != nil {
				return fmt.Errorf("cannot decode %v: %w", key, e)
			}
		case "meta":
			if e := dec.Decode(&ca.Meta); e != nil {
				return fmt.Errorf("cannot decode %v: %w", key, e)
			}
		case "provider":
			if e := dec.Decode(&ca.Provider); e != nil {
				return fmt.Errorf("cannot decode %v: %w", key, e)
			}
		default:
			var discard json.RawMessage
			if e := dec.Decode(&discard); e != nil {
				return fmt.Errorf("failed to skip unknown field %s: %w", key, e)
			}
		}
	}

	return nil
}

// Update client app information.
//
//	NOTE: Never update the login ID nor the user agent meta, these are only set
//	during registration.
func (ca *ClientApp) Update(store storage.Storage) error {
	ca.LastActivity = time.Now()

	// Store that token away for safe keeping
	if e := save(ca, store); e != nil {
		return e
	}

	return nil
}

// clientAppLocation containing login information.
func clientAppLocation(id string) string {
	return prefixClientApp + id + ".json"
}

// save client app information to storage.
func save(ca *ClientApp, store storage.Storage) error {
	data, e1 := json.Marshal(ca)
	if e1 != nil {
		return fmt.Errorf(stderr.EncodeJSON, e1.Error())
	}

	return store.Save(clientAppLocation(ca.id), data)
}
