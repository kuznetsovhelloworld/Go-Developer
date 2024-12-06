package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// menu prints out the available options for the user.
// This function is called repeatedly to show the main menu.
func menu() {
	fmt.Println("==== User Authentication System ====")
	fmt.Println("1) Register")
	fmt.Println("2) Login")
	fmt.Println("3) List All Users (Admin View)")
	fmt.Println("4) Quit")
	fmt.Print("Enter choice: ")
}

func main() {
	// Attempt to load existing users from the "users.json" file.
	// If the file doesn’t exist, an empty slice of users is returned.
	users, err := LoadUsers("users.json")
	if err != nil {
		// If there's an error loading users, print it out and proceed with an empty list.
		fmt.Println("Error loading users:", err)
		users = []User{}
	}

	// Create a new scanner to read input from the terminal.
	scanner := bufio.NewScanner(os.Stdin)

	// This loop keeps the program running until the user chooses to quit (option 4).
	for {
		// Display the menu each time the loop starts.
		menu()

		// Read the user’s choice.
		if !scanner.Scan() {
			// If scanning fails (e.g., EOF), break out of the loop and end the program.
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		// Use a switch statement to handle the different user actions.
		switch choice {
		case "1": // Register a new user
			fmt.Print("Enter new username: ")
			if !scanner.Scan() {
				// If we can't read input, break back to the menu.
				break
			}
			username := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter new password: ")
			if !scanner.Scan() {
				break
			}
			password := strings.TrimSpace(scanner.Text())

			// Check if the username is already taken.
			if userExists(users, username) {
				fmt.Println("User already exists. Please choose a different username.")
				// 'continue' sends us back to the start of the loop without executing more code in this case block.
				continue
			}

			// Register the user by creating a new User instance with a hashed password.
			newUser, err := RegisterUser(username, password)
			if err != nil {
				fmt.Println("Error registering user:", err)
				continue
			}

			// Add the new user to our list of users.
			users = append(users, newUser)

			// Save the updated list of users in a separate goroutine to demonstrate concurrency.
			go func(u []User) {
				err := SaveUsers("users.json", u)
				if err != nil {
					fmt.Println("Error saving users:", err)
				}
			}(users)

			fmt.Println("User registered successfully!")
			fmt.Println()

		case "2": // Login with existing credentials
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

			// Attempt to log in the user.
			ok := LoginUser(users, username, password)
			if ok {
				fmt.Println("Login successful!")
			} else {
				fmt.Println("Login failed. Check username and password.")
			}
			fmt.Println()

		case "3": // List all registered users (for demonstration purposes, acting as an "admin view")
			fmt.Println("Listing all users:")
			for _, u := range users {
				// Print each user's username.
				fmt.Printf(" - %s\n", u.Username)
			}
			fmt.Println()

		case "4": // Quit the application
			fmt.Println("Exiting...")
			// Sleep for a moment to allow the saving goroutine to finish if it is still running.
			time.Sleep(1 * time.Second)
			return

		default:
			// If user enters an invalid choice, prompt them to try again.
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
