package auth

import (
	"fmt"
)

type UserAuth struct {
	Users map[string]*User
}

func (ua *UserAuth) GetUser(username string) (*User, error) {
	u, ok := ua.Users[username]
	if !ok {
		return nil, fmt.Errorf("unknown user: %s", username)
	}

	return u, nil
}

func (ua *UserAuth) AddUser(username, password string) error {
	_, ok := ua.Users[username]
	if ok {
		return fmt.Errorf("username %s already exists", username)
	}

	u := NewUser(password)
	ua.Users[username] = u

	return nil
}
