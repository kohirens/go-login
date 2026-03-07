package login

type Profile struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	UserInfo *UserInfo `json:"userInfo"`
}

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
