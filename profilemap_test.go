package login

import (
	"os"
	"testing"

	"github.com/kohirens/storage"
)

func TestProfileMap(t *testing.T) {
	_ = os.MkdirAll(tmpDir+"/"+prefixProfileMap, os.ModePerm)
	fixedStore, e1 := storage.NewLocalStorage(tmpDir)
	if e1 != nil {
		t.Fatal(e1)
		return
	}

	cases := []struct {
		name    string
		id      string
		store   storage.Storage
		want    string
		wantErr bool
	}{
		{
			"success",
			"email@example.com",
			fixedStore,
			"rando-profile-id",
			false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if e := SaveProfileMap(c.id, c.want, c.store); (e != nil) != c.wantErr {
				t.Fatalf("SaveProfileMap() error = %v, wantErr %v", e, c.wantErr)
				return
			}
			got, err := LoadProfileMap(c.id, c.store)
			if (err != nil) != c.wantErr {
				t.Errorf("LoadProfileMap() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if got != c.want {
				t.Errorf("LoadProfileMap() got = %v, want %v", got, c.want)
				return
			}

			if e := DeleteProfileMap(c.id, c.store); (e != nil) != c.wantErr {
				t.Errorf("DeleteProfileMap() error = %v, wantErr %v", e, c.wantErr)
				return
			}
			// Make sure entry has been deleted.
			_, e1 = LoadProfileMap(c.id, c.store)
			if e1 == nil {
				t.Errorf("LoadProfileMap() error = %v, wantErr %v", err, c.wantErr)
				return
			}
		})
	}
}
