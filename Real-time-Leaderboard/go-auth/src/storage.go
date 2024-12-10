package main

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// fileMutex ensures that only one goroutine can write to the file at any given time.
// This prevents data corruption that could happen if multiple goroutines try to
// write to the same file simultaneously.
var fileMutex sync.Mutex

// LoadUsers attempts to load user data from the specified JSON file.
//
// Steps:
// 1. Check if the file exists. If it doesn't, return an empty user list with no error.
// 2. If the file does exist, open it and decode the JSON data into a []User slice.
// 3. Return the slice of users or any decoding error encountered.
func LoadUsers(filename string) ([]User, error) {
	// Check if the file exists. If not, return an empty slice of users.
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return []User{}, nil
	}

	// Open the file for reading.
	file, err := os.Open(filename)
	if err != nil {
		// If the file can't be opened (permissions, locked, etc.), return the error.
		return nil, err
	}
	defer file.Close() // Ensure the file is properly closed when the function exits.

	var users []User

	// Decode the JSON-encoded array of users into our users slice.
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		// If decoding fails (malformed JSON, unexpected structure), return the error.
		return nil, err
	}

	// Successfully loaded users, return them.
	return users, nil
}

// SaveUsers writes the given slice of users to the specified JSON file.
//
// Steps:
// 1. Acquire the mutex lock to ensure that only one goroutine writes at a time.
// 2. Create (or overwrite) the given file.
// 3. Encode the users slice into JSON and write it to the file in a readable, indented format.
// 4. Release the mutex lock and return any errors that occur during file creation or encoding.
func SaveUsers(filename string, users []User) error {
	// Lock the fileMutex to prevent concurrent writes.
	fileMutex.Lock()
	defer fileMutex.Unlock()

	// Create or overwrite the file. If a file with the same name exists, it will be replaced.
	file, err := os.Create(filename)
	if err != nil {
		// If the file cannot be created or opened for writing, return an error.
		return err
	}
	defer file.Close() // Ensure the file is closed after we're done writing.

	// Create a JSON encoder to write directly to the file.
	encoder := json.NewEncoder(file)
	// SetIndent makes the JSON file more human-readable by formatting it with spaces and newlines.
	encoder.SetIndent("", "  ")

	// Encode the users slice into JSON and write it to the file.
	return encoder.Encode(users)
}
