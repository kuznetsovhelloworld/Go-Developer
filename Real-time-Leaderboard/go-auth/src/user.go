package main

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user account in the system.
// Fields:
// - Username: the unique identifier chosen by the user.
// - Password: a bcrypt-hashed password string, not stored in plaintext.
// - LastLogin: a timestamp indicating when the user last successfully logged in.
type User struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"last_login"`
}

// RegisterUser validates a username and password, checks password complexity,
// and creates a new User with a securely hashed password.
// Returns a User struct and an error if any validation fails.
//
// Steps:
// 1. Ensure username and password are not empty.
// 2. Check that the password meets complexity requirements (length, digit, special character).
// 3. Hash the password using bcrypt.
// 4. Return a newly created User with an empty LastLogin time.
func RegisterUser(username, password string) (User, error) {
	if username == "" || password == "" {
		return User{}, errors.New("username or password cannot be empty")
	}

	// Validate password complexity before proceeding.
	if err := CheckPasswordComplexity(password); err != nil {
		return User{}, err
	}

	// Hash the password using bcrypt for secure storage.
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	// Return the new User with the hashed password and no last login time set yet.
	return User{
		Username:  username,
		Password:  string(hashed),
		LastLogin: time.Time{}, // User has not logged in yet, so this is zero-value time.
	}, nil
}

// LoginUser attempts to authenticate a user by comparing the provided username
// and plaintext password against stored credentials in the `users` slice.
//
// Steps:
// 1. Iterate over the list of known users.
// 2. If a matching username is found, compare the provided password with the stored hashed password.
// 3. If the password matches, update the user's LastLogin time to the current time.
// 4. Return true if a successful login occurs, otherwise false if authentication fails.
//
// Note: After successful login, you'd typically re-save the updated user list
// so that the changed LastLogin time persists.
func LoginUser(users []User, username, password string) bool {
	for i, u := range users {
		if u.Username == username {
			// Compare the provided password to the stored hashed password.
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
			if err == nil {
				// Password matches, update the last login time.
				users[i].LastLogin = time.Now()
				return true
			}
		}
	}
	return false
}

// userExists checks if a given username is already registered in the system.
// This prevents duplicate accounts.
//
// Steps:
// 1. Iterate through the users slice.
// 2. If any user has the same username, return true.
// 3. If no match is found, return false.
func userExists(users []User, username string) bool {
	for _, u := range users {
		if u.Username == username {
			return true
		}
	}
	return false
}
