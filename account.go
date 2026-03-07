package login

type Account struct {
	Profiles []*SubProfile `json:"profiles"`
}

type SubProfile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
