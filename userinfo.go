package login

// UserInfo is the personally identifiable information of a client/users
// profile. It MUST be kept secure at all times when it is not required for
// processing.
type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func NewUserInfo(email, firstName, lastName, phone string) *UserInfo {
	u := &UserInfo{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Id:        generateId(),
	}

	validateUserInfo(u)

	return u
}

func validateUserInfo(info *UserInfo) {
	// TODO: Validate email, firstname, lastname, phone with better standards.
	if info == nil {
		panic(stderr.NoInfo)
	}
	if info.FirstName == "" && len(info.FirstName) < 100 {
		panic(stderr.UserFirstName)
	}
	if info.LastName == "" && len(info.FirstName) < 100 {
		panic(stderr.UserLastName)
	}
	if info.Phone == "" && len(info.FirstName) < 100 {
		panic(stderr.UserPhone)
	}
	if info.Email == "" {
		panic(stderr.UserEmail)
	}
}
