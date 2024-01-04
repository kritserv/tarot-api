package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"math/rand"
	"strconv"
	"time"

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
	router.HandleFunc("/api/v1/cards/random", func(w http.ResponseWriter, r *http.Request) {
			r = mux.SetURLVars(r, map[string]string{"n": "1"})
			RandomCards(w, r)
		})
	router.HandleFunc("/api/v1/cards/{name_short}", CardShow)
	router.HandleFunc("/api/v1/cards/search/q={q}", CardSearch)
	router.HandleFunc("/api/v1/cards/search/meaning={meaning}", CardSearch)
	router.HandleFunc("/api/v1/cards/search/meaning_rev={meaning_rev}", CardSearch)
	router.HandleFunc("/api/v1/cards/random/n={n:[0-9]+}", RandomCards)
	
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

func CardSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	meaning := strings.ToLower(vars["meaning"])
	q := strings.ToLower(vars["q"])
	meaning_rev := strings.ToLower(vars["meaning_rev"])

	var results []Card

	for _, card := range cards.Cards {
		if meaning != "" && (strings.Contains(strings.ToLower(card.MeaningUp), meaning) || strings.Contains(strings.ToLower(card.MeaningRev), meaning)) {
			results = append(results, card)
		} else if q != "" && (strings.Contains(strings.ToLower(card.Name), q) || strings.Contains(strings.ToLower(card.MeaningRev), q)) {
			results = append(results, card)
		} else if meaning_rev != "" && strings.Contains(strings.ToLower(card.MeaningRev), meaning_rev) {
			results = append(results, card)
		}
	}

	if len(results) == 0 {
		http.Error(w, "Card not found", http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(results)
	}
}

func RandomCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nStr := vars["n"]

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 || n > len(cards.Cards) {
		http.Error(w, "Invalid value for n", http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards.Cards), func(i, j int) { cards.Cards[i], cards.Cards[j] = cards.Cards[j], cards.Cards[i] })

	json.NewEncoder(w).Encode(cards.Cards[:n])
}
