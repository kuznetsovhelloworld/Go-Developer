package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// menu prints the main menu options for user interaction.
// It will be called repeatedly in the main loop to prompt the user for actions.
func menu() {
	fmt.Println("==== User Authentication System ====")
	fmt.Println("1) Register")
	fmt.Println("2) Login")
	fmt.Println("3) List All Users (Admin View)")
	fmt.Println("4) Quit")
	fmt.Print("Enter choice: ")
}

func main() {
	// Attempt to load existing users from "users.json".
	// If the file doesn't exist or is empty, it returns an empty slice.
	users, err := LoadUsers("users.json")
	if err != nil {
		fmt.Println("Error loading users:", err)
		users = []User{} // Fallback to an empty user list if loading fails.
	}

	// Create a new scanner to read user input from standard input (the terminal).
	scanner := bufio.NewScanner(os.Stdin)

	// Main loop: Continues until the user chooses option 4 (Quit).
	for {
		// Show the menu each iteration so the user can pick an action.
		menu()

		// Read the user's menu choice.
		if !scanner.Scan() {
			// If we cannot read from input (e.g., EOF), break out of the loop and end.
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1": // Register a new user
			fmt.Print("Enter new username: ")
			if !scanner.Scan() {
				// If we fail to read input for username, break back to the main menu loop.
				break
			}
			username := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter new password: ")
			if !scanner.Scan() {
				// If we fail to read input for password, break back to the main menu loop.
				break
			}
			password := strings.TrimSpace(scanner.Text())

			// Check if this username is already taken.
			if userExists(users, username) {
				fmt.Println("User already exists. Please choose a different username.")
				continue // Skip the rest of this case and go back to showing the menu.
			}

			// Attempt to create a new user with the given username and password.
			newUser, err := RegisterUser(username, password)
			if err != nil {
				// If password doesn't meet complexity or other errors occur, inform the user.
				fmt.Println("Error registering user:", err)
				continue
			}

			// Append the new user to our in-memory list of users.
			users = append(users, newUser)

			// Save the updated user list in a separate goroutine to avoid blocking.
			// This demonstrates concurrency in Go (a unique aspect of the language).
			go func(u []User) {
				err := SaveUsers("users.json", u)
				if err != nil {
					fmt.Println("Error saving users:", err)
				}
			}(users)

			fmt.Println("User registered successfully!")
			fmt.Println()

		case "2": // Login with an existing user account
			fmt.Print("Enter username: ")
			if !scanner.Scan() {
				break
			}
			username := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter password: ")
			if !scanner.Scan() {
				break
			}
			password := strings.TrimSpace(scanner.Text())

			// Attempt to log in using the provided credentials.
			ok := LoginUser(users, username, password)
			if ok {
				fmt.Println("Login successful!")

				// When a user successfully logs in, their LastLogin time is updated.
				// Now we should save the updated user list so that the next time we run the app,
				// the last login time is preserved.
				go func(u []User) {
					err := SaveUsers("users.json", u)
					if err != nil {
						fmt.Println("Error saving users:", err)
					}
				}(users)

			} else {
				fmt.Println("Login failed. Check username and password.")
			}
			fmt.Println()

		case "3": // List all registered users (Admin view)
			fmt.Println("Listing all users:")
			for _, u := range users {
				// If the user has never logged in, show "Never logged in".
				// Otherwise, show the formatted last login time.
				lastLoginStr := "Never logged in"
				if !u.LastLogin.IsZero() {
					lastLoginStr = u.LastLogin.Format(time.RFC1123)
				}
				fmt.Printf(" - %s | Last Login: %s\n", u.Username, lastLoginStr)
			}
			fmt.Println()

		case "4": // Quit the application
			fmt.Println("Exiting...")
			// Wait for 1 second to allow any background goroutines (like saving) to finish.
			time.Sleep(1 * time.Second)
			return // Exit the main function, thus ending the program.

		default:
			// If the user enters a choice that isn't 1, 2, 3, or 4, prompt them again.
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
