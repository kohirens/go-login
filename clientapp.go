package login

import (
	"encoding/json"
	"fmt"
	"time"
)

// ClientApp represents the client's application used to access this application
// such as a browser or other application that can view web applications.
type ClientApp struct {
	Id        string    `json:"id"`
	AccountId string    `json:"accountId"`
	LastDate  time.Time `json:"lastDate"`
}

func LoadClientApp(ec []byte) (*ClientApp, error) {
	if len(ec) == 0 {
		return NewClientApp(), nil
	}

	var ca *ClientApp
	if e1 := json.Unmarshal(ec, &ca); e1 != nil {
		return nil, fmt.Errorf(stderr.DecodeJSON, e1.Error())
	}

	return ca, nil
}

func NewClientApp() *ClientApp {
	return &ClientApp{
		Id:       generateId(),
		LastDate: time.Now().UTC(),
	}
}
