package auth

type User struct {
	Password string
}

// TODO of fucking course, at least hash the password
func NewUser(password string) *User {
	u := User{
		Password: password,
	}

	return &u
}

// TODO of fucking course, change it to hash check before merging
func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
