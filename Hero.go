package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/subosito/gotenv"
)

// HeroResults struct
type HeroResults struct {
	Result struct {
		Heroes []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"heroes"`
		Status int `json:"status"`
		Count  int `json:"count"`
	} `json:"result"`
}

func init() {
	gotenv.Load()
}

func main() {
	safeAPIKey := url.QueryEscape(os.Getenv("STEAM_API_KEY"))
	url := fmt.Sprintf("https://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001/?key=%s", safeAPIKey)

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	var record HeroResults

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("First hero name = ", record.Result.Heroes[0].Name)
}
