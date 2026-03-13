package login

import (
	"os"
	"testing"

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

			gotLogin, e2 := LoadLogin(c.email, c.password, c.store)
			if (e2 != nil) != c.wantErr {
				t.Errorf("LoadLogin() error = %v, wantErr %v", e2, c.wantErr)
				return
			}

			if gotLogin.Email != got.Email {
				t.Errorf("Save() error emials do not match, got = %v\n\twant %v\n", gotLogin.Email, got.Email)
				return
			}

			if gotLogin.Password != got.Password {
				t.Errorf("Save() error password hashes do not match, got = %v\n\twant %v\n", gotLogin.Password, got.Password)
				return
			}

			// delete login
			if e := DeleteLogin(got.Email, fixedStore); e != nil {
				t.Errorf("DeleteLogin() error = %v, wantErr %v", e, c.wantErr)
				return
			}

			// verify delete
			_, e3 := LoadLogin(c.email, c.password, c.store)
			if e3 == nil {
				t.Errorf("DeleteLogin() may have failed")
				return
			}

		})
	}
}
