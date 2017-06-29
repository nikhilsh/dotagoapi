package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Hero struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var heroes []Hero

func GetHeroEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range heroes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Hero{})
}

func main() {
	fmt.Println("vim-go")
}
