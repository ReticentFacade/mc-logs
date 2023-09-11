// main.go
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables from .env file:
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	username := os.Getenv("ATERNOS_USERNAME")
	password := os.Getenv("ATERNOS_PASSWORD")

	// Aternos login URL:
	loginURL := "https://aternos.org/go"

	// Creating an HTTP client:
	client := &http.Client{}

	// Login by sending POST req using above credentials:
	data := url.Values{}
	data.Add("user", username)
	data.Add("password", password)
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	// Login request:
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error logging in:", err)
		return
	}
	defer resp.Body.Close()

	// Checking HTTP status code (successful response = (status code) 200 OK):
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Login failed. Status code:", resp.Status)
		return
	}

	// Parse HTML of logged-in page (using goquery):
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error parsing logged-in page:", err)
		return
	}

	// Does any specific element indicate success?
	if doc.Find("#logged-in-indicator").Length() == 0 {
		fmt.Println("Login failed. Missing indicator element.")
		return
	}

	// Login successful:
	fmt.Println("Login successful!")
}
