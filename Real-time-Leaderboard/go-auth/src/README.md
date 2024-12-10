# User Authentication System

This project is a simple user authentication system written in Go. It allows users to register, login, and list all users (admin view). The user data is stored in a JSON file.

## Features

- **Register**: Allows a new user to register by providing a username and password.
- **Login**: Allows an existing user to login by providing their username and password.
- **List All Users (Admin View)**: Allows an admin to view all registered users.
- **Quit**: Exits the application.

## Code Overview

### `main.go`

The `main.go` file contains the main logic of the application.

#### Imports

```go
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)
```
- **bufio**: Provides buffered I/O operations.
- **fmt**: Implements formatted I/O.
- **os**: Provides a platform-independent interface to operating system functionality.
- **strings**: Implements simple functions to manipulate UTF-8 encoded strings.
- **time**: Provides functionality for measuring and displaying time.

#### Functions

- **menu()**: Prints the main menu options for user interaction. It will be called repeatedly in the main loop to prompt the user for actions.

```go
func menu() {
	fmt.Println("==== User Authentication System ====")
	fmt.Println("1) Register")
	fmt.Println("2) Login")
	fmt.Println("3) List All Users (Admin View)")
	fmt.Println("4) Quit")
	fmt.Print("Enter choice: ")
}
```

- **main()**: The entry point of the application. It attempts to load existing users from `users.json`. If the file doesn't exist or is empty, it initializes an empty user list. It then enters a main loop where it repeatedly shows the menu and processes user input.

```go
func main() {
	users, err := LoadUsers("users.json")
	if err != nil {
		fmt.Println("Error loading users:", err)
		users = []User{}
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		menu()

		if !scanner.Scan() {
			break
		}
		// Additional code to handle user input will go here.
	}
}
```

### Running the Application

1. **Install Go**: Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/).
2. **Clone the Repository**: Clone this repository to your local machine.
3. **Run the Application**: Navigate to the project directory and run the application using the following command:

```sh
go run main.go
```

# Storage

This part of the User Authentication System handles the loading and saving of user data to a JSON file. It ensures that user data is safely read from and written to the file, preventing data corruption through concurrent access.

## Code Overview

### `storage.go`

The `storage.go` file contains functions for loading and saving user data.

#### Imports

```go
import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)
```

- **encoding/json**: Provides functions to encode and decode JSON data.
- **errors**: Implements functions to manipulate errors.
- **sync**: Provides basic synchronization primitives such as mutual exclusion locks.

#### Variables

- **fileMutex**: A `sync.Mutex` to ensure that only one goroutine can write to the file at any given time. This prevents data corruption that could happen if multiple goroutines try to write to the same file simultaneously.

```go
var fileMutex sync.Mutex
```

#### Functions

- **LoadUsers(filename string) ([]User, error)**: Attempts to load user data from the specified JSON file.

```go
func LoadUsers(filename string) ([]User, error) {
	// Check if the file exists. If not, return an empty slice of users.
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return []User{}, nil
	}

	// Open the file for reading.
	file, err := os.Open(filename)
	if

 err

 != nil {
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

	return users, nil
}
```

##### Steps:

1. **Check if the file exists**: If it doesn't, return an empty user list with no error.
2. **Open the file**: If the file exists, open it for reading.
3. **Decode the JSON data**: Decode the JSON-encoded array of users into a slice of `User`.
4. **Return the result**: Return the slice of users or any decoding error encountered.

# User Management

This part of the User Authentication System handles user account management, including user registration and password hashing.

## Code Overview

### `user.go`

The `user.go` file contains the `User` struct and functions for user registration and password management.

#### Imports

```go
import (
	"errors"
	"time"
	"golang.org/x/crypto/bcrypt"
)
```
- **time**: Provides functionality for measuring and displaying time.
- **golang.org/x/crypto/bcrypt**: Provides functions for hashing and comparing passwords using bcrypt.

#### Structs

- **User**: Represents a user account in the system.

```go
type User struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"last_login"`
}
```

- **Fields**:
  - `Username`: The unique identifier chosen by the user.
  - `Password`: A bcrypt-hashed password string, not stored in plaintext.
  - `LastLogin`: A timestamp indicating when the user last successfully logged in.

#### Functions

- **RegisterUser(username, password string) (User, error)**: Validates a username and password, checks password complexity, and creates a new `User` with a securely hashed password.

```go
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

	// Return a newly created User with an empty LastLogin time.
	return User{
		Username:  username,
		Password:  string(hashed),
		LastLogin: time.Time{},
	}, nil
}
```

##### Steps:

1. **Ensure username and password are not empty**: Returns an error if either is empty.
2. **Check password complexity**: Ensures the password meets complexity requirements (length, digit, special character).
3. **Hash the password**: Uses bcrypt to hash the password for secure storage.
4. **Return the new User**: Returns a newly created `User` with an empty `LastLogin` time.


# Utility Functions
This part of the User Authentication System includes utility functions that support the main functionality, such as validating password complexity.

## Code Overview
### `util.go`
The `util.go` file contains functions for validating password complexity.

#### Imports
```go
import (
    "errors"
    "unicode"
)
```
- **errors**: Implements functions to manipulate errors.
- **unicode**: Provides functions to test properties of Unicode code points.

#### Functions
- **CheckPasswordComplexity(password string) error**: Validates a password string against specific criteria to enforce stronger password rules and enhance security.
```go
func CheckPasswordComplexity(password string) error {
    // Check length first to ensure password isn't too short.
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters long")
    }

    // Track the presence of required character types.
    hasDigit := false
    hasSpecial := false

    // Iterate through each character in the password to identify digits and special characters.
    for _, c := range password {
        switch {
        case unicode.IsDigit(c):
            // If character is a digit, mark that we found a digit.
            hasDigit = true
        case unicode.IsSymbol(c) || unicode.IsPunct(c):
            // If character is a symbol or punctuation, mark that we found a special character.
            hasSpecial = true
        }
    }

    // After checking all characters, ensure that both digit and special character requirements are met.
    if !hasDigit {
        return errors.New("password must contain at least one digit")
    }
    if !hasSpecial {
        return errors.New("password must contain at least one special character")
    }

    return nil
}
```
#### Criteria Enforced:
1. **Minimum length of 8 characters**: Ensures the password isn't too short.
2. **Contains at least one digit**: Ensures the password includes numeric characters.
3. **Contains at least one special character**: Ensures the password includes symbols or punctuation.