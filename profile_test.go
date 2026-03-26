package login

import (
	"os"
	"reflect"
	"testing"

	"github.com/kohirens/storage"
)

const (
	tmpDir = "tmp"
)

func TestMain(m *testing.M) {
	_ = os.MkdirAll(tmpDir, 0777)
	os.Exit(m.Run())
}

func TestProfile_ClientApp(t *testing.T) {
	cases := []struct {
		name        string
		clientAppId *ClientApp
		profile     *Profile
		wantErr     bool
	}{
		{
			"success_add_find_delete_client_app",
			&ClientApp{Id: "test-01-client-app-id-01"},
			&Profile{
				ClientApp: map[string]*ClientApp{},
				Id:        "test-01-profile-clientapp-id",
				Name:      "test-01-profile-clientapp-name",
			},
			false,
		},
	}
	for _, c := range cases {
		// Can add a ClientApp to a profile.
		c.profile.AddClientApp(c.clientAppId)

		// Can find a ClientApp in a profile.
		got, e2 := c.profile.FindClientApp(c.clientAppId.Id)

		if (e2 != nil) != c.wantErr {
			t.Errorf("Profile.FindClientApp(%v) gotErr %v, wantErr %v", c.clientAppId, e2, c.wantErr)
			return
		}

		if !reflect.DeepEqual(got, c.clientAppId) {
			t.Errorf("FindClientApp() got = %v, want %v", got, c.clientAppId)
			return
		}

		// Can delete a ClientApp from a profile.
		if e := c.profile.RemoveClientApp(c.clientAppId.Id); (e != nil) != c.wantErr {
			t.Errorf("Profile.RemoveClientApp(%v) got err %v", c.clientAppId.Id, e)
			return
		}
	}
}

func TestProfile_Save(t *testing.T) {
	_ = os.MkdirAll(tmpDir+"/"+prefixProfile, os.ModePerm)
	fixedStore, e1 := storage.NewLocalStorage(tmpDir)
	if e1 != nil {
		panic(e1)
	}

	cases := []struct {
		name        string
		ClientAppId string
		UserInfo    *UserInfo
		store       storage.Storage
		wantErr     bool
	}{
		{
			"manual_account",
			"manual-client-app-id-01",
			&UserInfo{
				Id:        "manual-user-id",
				FirstName: "SaveTest-01",
				LastName:  "SaveLastTest-02",
				Email:     "test-01@example.com",
				Phone:     "555-555-5555",
			},
			fixedStore,
			false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := NewProfile(c.name, generateId(), c.UserInfo)
			// Save the profile
			if err := p.Save(c.store); (err != nil) != c.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, c.wantErr)
				return
			}

			// Read the profile
			gotProf, e2 := LoadProfile(p.Id, fixedStore)
			if e2 != nil {
				t.Errorf("LoadProfile() error = %v", e2.Error())
				return
			}

			if gotProf.UserInfo.Email != p.UserInfo.Email {
				t.Errorf("Save() error user emails do not match, got = %v\n\twant %v\n", gotProf.UserInfo.Email, p.UserInfo.Email)
				return
			}
			if gotProf.Name != p.Name {
				t.Errorf("LoadProfile() error Name do not match, got = %v\n\twant %v\n", gotProf.Name, p.Name)
				return
			}
			if gotProf.Id != p.Id {
				t.Errorf("LoadProfile() error ID do not match, got = %v\n\twant %v\n", gotProf.Id, p.Id)
				return
			}

			if e := DeleteProfile(p.Id, fixedStore); e != nil {
				t.Errorf("DeleteProfile() error = %v", e.Error())
				return
			}

			// Read the profile
			_, e3 := LoadProfile(p.Id, fixedStore)
			if e3 == nil {
				t.Errorf("DeleteProfile may not have been successful")
				return
			}
		})
	}
}
