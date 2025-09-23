package auth

import "fmt"

type UserAuth struct {
	Users map[string]User
}

func (ua *UserAuth) GetUser(username string) (*User, error) {
	u, ok := ua.Users[username]
	if !ok {
		return nil, fmt.Errorf("unknown user: %s", username)
	}

	return &u, nil
}
