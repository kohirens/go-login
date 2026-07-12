package login

import (
	"testing"
	"time"

	"github.com/kohirens/sso/oidc"
	"github.com/mileusna/useragent"
)

func TestClientApp_String(t *testing.T) {
	cases := []struct {
		name         string
		id           string
		LastActivity time.Time
		Provider     oidc.Provider
		Meta         *useragent.UserAgent
		want         string
	}{
		{
			name:         "success",
			id:           "123",
			LastActivity: time.Date(2026, 07, 11, 01, 18, 0, 0, time.UTC),
			Provider:     nil,
			Meta:         &useragent.UserAgent{},
			want:         `{"id": "123","lastActivity": "2026-07-11 01:18:00 +0000 UTC","meta": ""}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ca := &ClientApp{
				id:           c.id,
				LastActivity: c.LastActivity,
				Provider:     c.Provider,
				Meta:         c.Meta,
			}
			if got := ca.String(); got != c.want {
				t.Errorf("String() = %v, want %v", got, c.want)
			}
		})
	}
}
