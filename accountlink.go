package login

// AccountLink link an account by email and password.
type AccountLink struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	AccountID string `json:"accountID"`
}
