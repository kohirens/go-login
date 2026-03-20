package login

import (
	"time"
)

// ClientApp represents the client's application used to access this application
// such as a browser or other application that can view web applications.
type ClientApp struct {
	Id        string    `json:"id"`
	UserAgent string    `json:"userAgent"`
	AccountId string    `json:"accountId"`
	LastDate  time.Time `json:"lastDate"`
}

func NewClientApp(userAgent string) *ClientApp {
	return &ClientApp{
		Id:        generateId(),
		UserAgent: userAgent,
		LastDate:  time.Now().UTC(),
	}
}
