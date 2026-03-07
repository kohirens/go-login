package login

import (
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
		name     string
		DeviceId string
		UserInfo *UserInfo
		store    storage.Storage
		wantErr  bool
	}{
		{
			"manual_account",
			"manual-device-id-01",
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
			p := NewProfile(generateId(), c.UserInfo)
			if err := p.Save(c.store); (err != nil) != c.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, c.wantErr)
			}
		})
	}
}
