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

func getHeroesList() HeroResults {
	heroesListURL := fmt.Sprintf("https://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001/?key=%s", url.QueryEscape(os.Getenv("STEAM_API_KEY")))
	req, err := http.NewRequest("GET", heroesListURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return HeroResults{}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return HeroResults{}
	}
	defer resp.Body.Close()

	var record HeroResults

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	return record
}

func getImageFor(heroName string) []byte {
	heroImageURL := fmt.Sprintf("http://cdn.dota2.com/apps/dota2/images/heroes/%s_sb.png", heroName)
	req, err := http.NewRequest("GET", heroImageURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return []byte{}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return []byte{}
	}
	defer resp.Body.Close()

	var record HeroResults

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	//save image to database with hero name
	return []byte{}
}

func main() {
}
