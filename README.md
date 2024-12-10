

# Go Programming Language Project

## Overview

This project explores the Go programming language, commonly known as Golang. Developed by Google, Go is designed for simplicity, efficiency, and reliability, making it ideal for building scalable and high-performance applications. This repository includes documentation on Go's features and a sample project demonstrating its capabilities.

## Purpose

Our primary goal is to share our experiences with the Go programming language and to learn from the community. By documenting our journey, we aim to provide valuable insights to others and welcome feedback to enhance our understanding.

## Table of Contents

- [Go Programming Language Project](#go-programming-language-project)
  - [Overview](#overview)
  - [Purpose](#purpose)
  - [Table of Contents](#table-of-contents)
  - [Introduction to Go](#introduction-to-go)
  - [Installation Guide](#installation-guide)
  - [Advanced Features](#advanced-features)
  - [Sample Project](#sample-project)
  - [Challenges and Solutions](#challenges-and-solutions)
  - [Learning Resources](#learning-resources)
  - [Contributors](#contributors)

## Introduction to Go

Go, often referred to as Golang, is an open-source programming language developed by Google. It was created to address the challenges of modern software development, emphasizing simplicity, concurrency, and efficient execution. Go is statically typed and compiles to native machine code, offering the performance benefits of languages like C++ with the readability and ease of use found in dynamically typed languages.

## Installation Guide

To set up Go on your system:

1. **Download**: Visit the official Go website at [https://go.dev/dl/](https://go.dev/dl/) and download the installer for your operating system.
2. **Install**: Run the installer and follow the on-screen instructions.
3. **Verify Installation**: Open a terminal or command prompt and execute:

   ```bash  
    go version
    ```  
This command should display the installed Go version, confirming a successful installation.

## Basic Syntax and Features

Go's syntax is designed to be clean and concise. Here's a simple "Hello, World!" program:

```go  
package main  
  
import "fmt"  
  
func main() {  
 fmt.Println("Hello, World!")}  
  ```  

Key features of Go include:

- **Variables and Data Types**: Go supports various data types such as `int`, `float64`, `string`, and `bool`. Variables can be declared using the `var` keyword or inferred using `:=`.

  ```go  
    var x int = 10  
    y := 20
  ```

- **Control Structures**: Go provides standard control structures like `if`, `for`, and `switch`.  
  
  ```go  
    if x > y {
    fmt.Println("x is greater than y")
    } else { fmt.Println("x is less than or equal to y")
    }  
  ```

- **Functions**: Functions are first-class citizens in Go and can return multiple values.

  ```go  
  func add(a int, b int) int {
  return a + b }
  ``` 

  - **Multiple Return Values**: Go has built-in support for multiple return values. This feature is often used in Go, for example to return result and error values ​​from a function.
    ```go
    func vals() (int, int) {
      return 3, 7
    }
    func main(){
      //Here we use the 2 different return values from the call with multiple assignment.
      a, b := vals()
      fmt.Println(a)
      fmt.Println(b)
    
      //If you only want a subset of the returned values, use the blank identifier _.
      _, c := vals()
      fmt.Println(c)
    }
    ```

  - **Variadic Functions**: accepts multiple arguments of the same type and can be called with any number of arguments, including none.
    ```go
    // Variadic function to calculate sum
    func sum(nums ...int) int {
      total := 0
      for _, n := range nums {
        total += n
      }
      return total
    }

    func main() {
      fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
      fmt.Println("Sum of 4, 5:", sum(4, 5))
      fmt.Println("Sum of no numbers:", sum())
   }
    ```

## Advanced Features  
  
Go offers several advanced features:  
  
- **Concurrency**: Go's concurrency model is based on goroutines and channels, allowing efficient execution of concurrent tasks.  
  
  ```go  
  go func() {
  fmt.Println("This runs concurrently") 
  }
  ``` 
  - **Multiple Goroutines**:You can create a goroutine simply by using go keyword as a prefixing to the function or method call as shown in the below syntax.
    ```go
    func f(from string) {
      for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
      }
    }
  
    func main() {
      f("direct")
  
      //To invoke this function in a goroutine, use go f(s).
      //This new goroutine will execute concurrently with the calling one.
      go f("goroutine")
    
      //You can also start a goroutine for an anonymous function call.
      go func(msg string) {
        fmt.Println(msg)
      }("going")
    
      //The two function calls are running asynchronously in separate goroutines now.

      //Wait for them to finish
      time.Sleep(time.Second)
      fmt.Println("done")
    }
    ```
  - **WaitGroups**: To wait for multiple goroutines to finish
    ```go
    var wg sync.WaitGroup
    //Launch several goroutines and increment the WaitGroup counter for each.
    for i := 1; i <= 5; i++ {
      wg.Add(1)
    
      go func() {
        defer wg.Done()
        worker(i)
      }()
    }
    //Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
    wg.Wait()
    ```

- **Structs and Interfaces**: Go uses structs to group related data and interfaces to define behavior.

  ```go
  type Person struct {
  Name string
  Age  int
  }
  type Greeter interface {
  Greet() string
  }
  ```

- **Error Handling**: Go emphasizes explicit error handling, typically returning error values from functions.  
  
  ```go
  func divide(a, b float64) (float64, error) {
  if b == 0 {
  return 0, errors.New("division by zero")
  }
  return a / b, nil
  }
  ```

## Sample Project

This repository includes a sample project that demonstrates Go's features. The project is a simple command-line application that performs basic CRUD operations. To run the project:

1. Navigate to the project directory. (where go.mod is located)
2. Execute:   `  go run ./src`
   Follow the on-screen instructions to register, log in, list users, or quit.

**Features Shown:**

-   User registration and login with bcrypt password hashing.
-   JSON file storage for persistent user data.
-   Concurrency demonstrated through goroutines that save data in the background.
-   Conditional logic, loops, and error handling that reflect Go’s idiomatic patterns.

## Challenges and Solutions

During the development of this project, we encountered several challenges:

- **Understanding Goroutines**: Managing concurrent processes required a thorough understanding of goroutines and channels. We solved this by reviewing Go's concurrency documentation and experimenting with simple examples before integrating them into the projec.

- **Error Handling**: Adjusting to explicit error returns rather than exceptions required us to be more diligent. We addressed this by systematically checking errors after every file operation, user registration, and login attempt, and by logging helpful error messages.

## Learning Resources

To further explore Go, consider the following resources:

- **Official Documentation**: Comprehensive guides and references are available at [https://go.dev/doc/](https://go.dev/doc/).
- **A Tour of Go**: An interactive tutorial to learn Go's basics, accessible at [https://tour.golang.org/](https://tour.golang.org/).
- **Go by Example**: A collection of annotated code examples demonstrating various Go concepts, found at [https://gobyexample.com/](https://gobyexample.com/).

## Contributors

- [Mykyta Kuznetsov](https://github.com/kuznetsovhelloworld)
- [Rodrigo Sanchez](https://github.com/rodrigosanchezq)

We welcome contributions and feedback. Please feel free to open issues or submit pull requests.