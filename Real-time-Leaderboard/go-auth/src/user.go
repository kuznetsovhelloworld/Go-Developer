package main

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterUser creates a new user with a hashed password
func RegisterUser(username, password string) (User, error) {
	if username == "" || password == "" {
		return User{}, errors.New("username or password cannot be empty")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	return User{
		Username: username,
		Password: string(hashed),
	}, nil
}

// LoginUser checks if the username and password match an existing user
func LoginUser(users []User, username, password string) bool {
	for _, u := range users {
		if u.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
			if err == nil {
				return true
			}
		}
	}
	return false
}

// userExists checks if a username is already taken
func userExists(users []User, username string) bool {
	for _, u := range users {
		if u.Username == username {
			return true
		}
	}
	return false
}
