package login

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/kohirens/storage"
)

func TestAccount_Save(t *testing.T) {
	_ = os.MkdirAll(tmpDir+"/"+prefixAccount, os.ModePerm)
	_ = os.MkdirAll(tmpDir+"/"+prefixProfile, os.ModePerm)
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
				ID:        "google-userid-00001",
				FirstName: "Test-01",
				LastName:  "Account-01",
				Email:     "test-01@example.com",
				Phone:     "555-555-5555",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProfile := NewProfile(tt.profile, tt.user)
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
			gotAct, e2 := LoadAccount(act.id, fixedStore)
			if e2 != nil {
				t.Fatal(e2)
				return
			}

			// Verify data integrity
			if act.id != gotAct.id {
				t.Errorf("Save() got = %v, want = %v", gotAct.id, act.id)
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
			if e := DeleteAccount(gotAct.id, fixedStore); e != nil {
				t.Fatal(e)
				return
			}

			// Make sure the account was deleted.
			_, e3 := LoadAccount(act.id, fixedStore)
			if e3 == nil {
				t.Fatal(fmt.Sprintf("cannot delete the account %v", gotAct.id))
				return
			}
		})
	}
}

func TestAccount_String(t *testing.T) {
	cases := []struct {
		name    string
		account *Account
		want    string
	}{
		{
			"stringify_success",
			&Account{
				Owner:    "1",
				Profiles: map[string]*SubProfile{"test": {"1", "test"}},
				id:       "1234",
			},
			`{"id":"1234","owner":"1","profiles":{"test":{"id":"1","name":"test"}}}`,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.account.String(); got != tt.want {
				t.Errorf("\nString():\n\t got := %v\n\twant := %v\n", got, tt.want)
			}
		})
	}
}

func TestAccount_Unmarshal(t *testing.T) {
	cases := []struct {
		name    string
		data    []byte
		want    *Account
		wantErr bool
	}{
		{
			"success",
			[]byte(`{"id":"1234","owner":"1234","profiles":{"test":{"id":"1","name":"test"}}}`),
			&Account{
				Owner:    "1234",
				Profiles: map[string]*SubProfile{"test": {"1", "test"}},
				id:       "1234",
			},
			false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := &Account{}
			if err := got.UnmarshalJSON(c.data); (err != nil) != c.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, c.wantErr)
			}

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("Unmarshal() got = %v, want %v", got, c.want)
			}
		})
	}
}
