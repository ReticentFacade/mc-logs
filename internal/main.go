// main.go
package main

import (
	"fmt"
	"log"
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
		log.Println("Error creating request:", err)
	}
	log.Println(`
	POST request created:
	URL:`, req.URL, `
	Body:`, req.Body, `
	Header:`, req.Header, `

	Username: `, username, `
	Password: `, password, `
	`)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	log.Println("Sending POST request...")
	log.Println("Request:", req)

	// Login request:
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Checking HTTP status code (successful response = (status code) 200 OK):
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: login failed. Status code:", resp.StatusCode)
	}

	// Parse HTML of logged-in page (using goquery):
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error parsing HTML:", err)
	}

	// Does any specific element indicate success?
	if doc.Find("#logged-in-indicator").Length() == 0 {
		log.Println("Error: login failed. No logged-in indicator found.")
	}

	// Login successful:
	log.Println("Login successful!")
}
