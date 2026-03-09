package login

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/kohirens/storage"
)

func TestLogin_Save(t *testing.T) {
	_ = os.MkdirAll(tmpDir+"/"+prefixLogin, os.ModePerm)
	fixedStore, e1 := storage.NewLocalStorage(tmpDir)
	if e1 != nil {
		t.Fatal(e1)
		return
	}

	cases := []struct {
		name     string
		email    string
		password string
		store    storage.Storage
		wantErr  bool
	}{
		{
			"success",
			"test+login@example.com",
			"1234",
			fixedStore,
			false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := NewLogin(c.email, c.password)

			if err := got.Save(c.store); (err != nil) != c.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, c.wantErr)
			}

			gotLoc := loginLocation(got.Email)
			gotData, _ := os.ReadFile(tmpDir + "/" + gotLoc)
			var want *Login
			if e := json.Unmarshal(gotData, &want); e != nil {
				t.Errorf("Unmarshall Save() error = %v", e.Error())
				return
			}

			if want.Email != got.Email {
				t.Errorf("Save() error emials do not match, got = %v\n\twant %v\n", got.Email, want.Email)
			}

			if want.Password != got.Password {
				t.Errorf("Save() error password hashes do not match, got = %v\n\twant %v\n", got.Password, want.Password)
			}
		})
	}
}

func TestLoadLogin(t *testing.T) {
	_ = os.MkdirAll(tmpDir+"/"+prefixLogin, os.ModePerm)
	fixedStore, e1 := storage.NewLocalStorage(tmpDir)
	if e1 != nil {
		t.Fatal(e1)
		return
	}

	cases := []struct {
		name     string
		email    string
		password string
		store    storage.Storage
		want     *Login
		wantErr  bool
	}{
		{
			"success",
			"test+login@example.com",
			"1234",
			fixedStore,
			&Login{
				Email:    "test+login@example.com",
				Password: "064b32c8-9099-52b9-a716-931a01d3f9f5",
				LastDate: time.Now().UTC(),
			},
			false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := LoadLogin(c.email, c.password, c.store)
			if (err != nil) != c.wantErr {
				t.Errorf("LoadLogin() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if c.want.Email != got.Email {
				t.Errorf("LoadLogin() email got = %v, want %v", got.Email, c.want.Email)
				return
			}

			if c.want.Password != got.Password {
				t.Errorf("LoadLogin() email got = %v, want %v", got.Password, c.want.Password)
				return
			}
		})
	}
}
