package login

type Account struct {
	Id       string                 `json:"id"`
	Owner    string                 `json:"owner"`
	Profiles map[string]*SubProfile `json:"profiles"`
}

type SubProfile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewAccount(profileId, profileName string) *Account {
	return &Account{
		Id: generateId(),
		Profiles: map[string]*SubProfile{
			profileId: {
				Id:   profileId,
				Name: profileName,
			},
		},
	}
}
