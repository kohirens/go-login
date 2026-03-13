package login

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func validateUserInfo(info *UserInfo) {
	// TODO: Validate email, firstname, lastname, phone with better standards.
	if info == nil {
		panic("nil UserInfo")
	}
	if info.FirstName == "" {
		panic("UserInfo FirstName is empty")
	}
	if info.LastName == "" {
		panic("UserInfo LastName is empty")
	}
	if info.Phone == "" {
		panic("UserInfo Phone is empty")
	}
	if info.Email == "" {
		panic("UserInfo Email is empty")
	}
}
