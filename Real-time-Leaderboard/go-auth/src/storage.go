package main

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// fileMutex is used to prevent concurrent write operations to the file.
// This ensures that if multiple goroutines try to save users at the same time,
// they won't interfere with each other, potentially causing data corruption.
var fileMutex sync.Mutex

// LoadUsers attempts to load user data from a given JSON file.
// If the file does not exist, it returns an empty slice of users without an error.
func LoadUsers(filename string) ([]User, error) {
	// Check if the file exists using os.Stat.
	// If it doesn't exist, return an empty slice (no error).
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return []User{}, nil
	}

	// Open the file in read-only mode.
	file, err := os.Open(filename)
	if err != nil {
		// If there is an error opening the file, return the error.
		return nil, err
	}
	defer file.Close() // Ensure the file is closed when the function returns.

	// Create a slice to hold the users that will be decoded from the file.
	var users []User

	// Decode the JSON data from the file into the users slice.
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		// If there's an error in decoding (malformed JSON, etc.), return it.
		return nil, err
	}

	// If everything is successful, return the slice of users.
	return users, nil
}

// SaveUsers writes the provided list of users to a JSON file.
// It uses a mutex to ensure only one write operation happens at a time.
func SaveUsers(filename string, users []User) error {
	// Lock the mutex before writing to the file to prevent concurrent writes.
	fileMutex.Lock()
	defer fileMutex.Unlock() // Unlock will be called once we return from this function.

	// Create (or overwrite) the file where user data will be stored.
	file, err := os.Create(filename)
	if err != nil {
		// If there's an error creating the file, return it.
		return err
	}
	defer file.Close() // Ensure the file is closed when we return.

	// Create a JSON encoder that writes directly to the file.
	encoder := json.NewEncoder(file)
	// SetIndent makes the JSON output more readable, adding pretty formatting.
	encoder.SetIndent("", "  ")

	// Encode the users slice into JSON and write it to the file.
	return encoder.Encode(users)
}
