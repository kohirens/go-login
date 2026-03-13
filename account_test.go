package login

import (
	"fmt"
	"os"
	"testing"

	"github.com/kohirens/storage"
)

func TestAccount_Save(t *testing.T) {
	_ = os.MkdirAll(tmpDir+"/"+prefixAccount, os.ModePerm)
	_ = os.MkdirAll(tmpDir+"/"+prefixProfile, os.ModePerm)
	_ = os.MkdirAll(tmpDir+"/"+prefixLogin, os.ModePerm)
	fixedStore, e1 := storage.NewLocalStorage(tmpDir)
	if e1 != nil {
		t.Fatal(e1)
		return
	}

	tests := []struct {
		name    string
		profile string
		appId   string
		user    *UserInfo
		wantErr bool
	}{
		{
			"success",
			"testAccountProfileName-01",
			"guid-test-0134",
			&UserInfo{
				Id:        "",
				FirstName: "Test-01",
				LastName:  "Account-01",
				Email:     "test-01@example.com",
				Phone:     "555-5555-5555",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProfile := NewProfile(tt.profile, tt.appId, tt.user)
			if e := gotProfile.Save(fixedStore); e != nil {
				t.Fatal(e)
				return
			}

			act := NewAccount(gotProfile.Id, gotProfile.Name)
			if err := act.Save(fixedStore); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Make sure the account was indeed saved.
			gotAct, e2 := LoadAccount(act.Id, fixedStore)
			if e2 != nil {
				t.Fatal(e2)
				return
			}

			// Verify data integrity
			if act.Id != gotAct.Id {
				t.Errorf("Save() got = %v, want = %v", gotAct.Id, act.Id)
			}
			if act.Owner != gotAct.Owner {
				t.Errorf("Save() got = %v, want = %v", gotAct.Owner, act.Owner)
			}
			for pId := range act.Profiles {
				if act.Profiles[pId].Name != gotAct.Profiles[pId].Name {
					t.Errorf("Save() got = %v, want = %v", gotAct.Profiles[pId].Name, act.Profiles[pId].Name)
				}
			}

			// Delete the account
			if e := DeleteAccount(gotAct.Id, fixedStore); e != nil {
				t.Fatal(e)
				return
			}

			// Make sure the account was deleted.
			_, e3 := LoadAccount(act.Id, fixedStore)
			if e3 == nil {
				t.Fatal(fmt.Sprintf("cannot delete the account %v", gotAct.Id))
				return
			}
		})
	}
}
