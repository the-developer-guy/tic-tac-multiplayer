package auth

type User struct {
	Password string
}

// TODO of fucking course, change it to hash check before merging
func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
