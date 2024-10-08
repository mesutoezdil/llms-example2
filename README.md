# llms-example2

This project demonstrates how to use the Prediction Guard API in a Go application to generate text completions. The goal was to interact with the Prediction Guard API to request Go code generation using a language model.

## Table of Contents

- Project Overview
- Setup
- Issues Faced
- Solution and Fixes
- Running the Program
- Generated Response Example
- Conclusion

## Project Overview

The goal of this project was to write a Go program that leverages the Prediction Guard API client to generate code based on input prompts. The initial code aimed to interact with the API and request a code snippet that prints random numbers.

## Setup

1. **Cloning the Repository**: This project was cloned from GitHub: https://github.com/mesutoezdil/llms-example2.git.

2. **Setting Up the Environment**: An environment variable for the API key was used as follows:  
   `export PGKEY="your_api_key_here"`  
   Replace `"your_api_key_here"` with your actual Prediction Guard API key.

3. **Go Modules Initialization**: The repository was set up with Go modules, and the necessary package `github.com/predictionguard/go-client` was installed at version `v0.21.0`:
   - `go mod init github.com/mesutoezdil/llms-example2`
   - `go get github.com/predictionguard/go-client@v0.21.0`

## Issues Faced

1. **Role Type Compatibility**: The initial issue was related to the `Role` type expected by the `go-client`. Earlier attempts to use plain strings for `"system"` and `"user"` roles resulted in errors as they didn't match the expected type.

2. **Version Compatibility**: The `go-client` library version `v0.21.0` required different handling of `Role` and `Model` than later versions, which led to confusion while trying to use roles and model strings correctly.

## Solution and Fixes

1. **Downgrading the `go-client` Version**: To resolve compatibility issues, the `go-client` package was downgraded to `v0.21.0`. This was necessary due to breaking changes in newer versions:
   - `go get github.com/predictionguard/go-client@v0.21.0`

2. **Correct Use of Roles**: After exploring the library documentation, it was found that the roles could be accessed using the `client.Roles` variable:
   - `client.Roles.System` for the system role
   - `client.Roles.User` for the user role

3. **Successful API Call**: After applying the correct `Role` and `Model` values, the API responded successfully with a code snippet as expected.

## Running the Program

1. **Run the Go Program**: After setting up the environment and fixing the compatibility issues, run the program using:
   - `go run main.go`
   - Ensure your environment variable `PGKEY` is set with your Prediction Guard API key.

## Generated Response Example

Upon successful execution, the program makes a request to the Prediction Guard API to generate a Go code snippet that prints out 10 random numbers. The API returned the following code example:

```
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 10; i++ {
        fmt.Println(rand.Intn(100))
    }
}
```

This code uses the `math/rand` package to generate random numbers and prints out 10 numbers between 0 and 99.

## Conclusion

The main challenges were understanding the `Role` type expectations of the Prediction Guard client library and ensuring the code was compatible with the `v0.21.0` version. By properly using `client.Roles` and downgrading the client library, the API was successfully accessed, and the desired output was generated.
