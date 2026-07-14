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

func TestDeleteClientApp(t *testing.T) {
	fixedStore, e1 := fixtureStore()
	if e1 != nil {
		t.Fatal(e1)
	}
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			"success",
			false,
		},
	}

	for _, c := range cases {
		fixtureClientApp, e2 := RegisterClientApp("", nil, fixedStore)
		if e2 != nil && !c.wantErr {
			t.Errorf("DeleteClientApp(%q) got error: %v", c.name, e2)
			return
		}
		t.Run(c.name, func(t *testing.T) {
			if err := DeleteClientApp(fixtureClientApp, fixedStore); (err != nil) != c.wantErr {
				t.Errorf("DeleteClientApp() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			// verify the file was in face deleted.
			got, e3 := LoadClientApp(fixtureClientApp.id, fixedStore)
			if got != nil || e3 == nil {
				t.Errorf("DeleteClientApp(%q) error: %v", c.name, e2)
				return
			}
		})
	}
}
