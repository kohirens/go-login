package login

import (
	"encoding/json"
	"os"
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

func TestProfile_Save(t *testing.T) {
	fixedStore, _ := storage.NewLocalStorage(tmpDir)
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
			if err := p.Save(c.store); (err != nil) != c.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, c.wantErr)
				return
			}

			gotData, _ := os.ReadFile(tmpDir + "/" + p.Id + ".json")
			var ep *Profile
			if e := json.Unmarshal(gotData, &ep); e != nil {
				t.Errorf("Unmarshall Save() error = %v", e.Error())
				return
			}

			if ep.Id != p.Id {
				t.Errorf("Save() error IDs do not match, got = %v\n\twant %v\n", ep.Id, p.Id)
			}

			if ep.UserInfo.Email != p.UserInfo.Email {
				t.Errorf("Save() error user emails do not match, got = %v\n\twant %v\n", ep.UserInfo.Email, p.UserInfo.Email)
			}
		})
	}
}
