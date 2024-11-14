// Go example with various secrets

package main

import "fmt"

func main() {
    // AWS Key
    awsAccessKey := "AKIAIOSFODNN7EXAMPLE"

    // API Key
    googleApiKey := "AIzaSyDqkBEXAMPLEEXAMPLEEXAMPLEEXAMPLE"

    // JWT Token
    jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.DuZbVfQtZ7sEXAMPLEEXAMPLEEXAMPLEEXAMPLE"

    // Password
    dbPassword := "123456789"

    // High-entropy GitHub token
    githubToken := "ghp_aBCdEFgHiJKlmNOPqrSTUVWXyZ12345"

    fmt.Println("Secrets loaded")
}
