package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Card struct {
	Type       string `json:"type"`
	NameShort  string `json:"name_short"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	ValueInt   int    `json:"value_int"`
	MeaningUp  string `json:"meaning_up"`
	MeaningRev string `json:"meaning_rev"`
	Desc       string `json:"desc"`
}

type Cards struct {
	Nhits int    `json:"nhits"`
	Cards []Card `json:"cards"`
}


var cards Cards

func init() {
	absPath, _ := filepath.Abs("static/card_data.json")
	file, err := ioutil.ReadFile(absPath)

	if err != nil {
		log.Fatalf("File not found: %v", err)
	}

	err = json.Unmarshal(file, &cards)

	if err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/api/v1/cards", CardsIndex)
	router.HandleFunc("/api/v1", CardsIndex)
	router.HandleFunc("/api/v1/cards/{name_short}", CardShow)
	
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")
	
	log.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func CardsIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cards)
}

func CardShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nameShort := vars["name_short"]

	for _, card := range cards.Cards {
		if card.NameShort == nameShort {
			json.NewEncoder(w).Encode(card)
			return
		}
	}

	http.Error(w, "Card not found", http.StatusNotFound)
}
