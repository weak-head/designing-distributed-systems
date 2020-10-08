package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/mux"
)

type state struct {
	mc *memcache.Client
}

func newState(mc *memcache.Client) *state {
	return &state{mc: mc}
}

func (s *state) handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	log.Printf("GET %s", key)

	value, err := s.mc.Get(key)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(value.Value))
}

func (s *state) handleSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	log.Printf("SET %s %s", key, value)

	s.mc.Set(&memcache.Item{Key: key, Value: []byte(value)})
	w.WriteHeader(http.StatusOK)
}

func main() {
	memHost := os.Getenv("MEMCACHE_HOST")
	memPort := os.Getenv("MEMCACHE_PORT")

	if memHost == "" {
		memHost = "localhost"
	}
	if memPort == "" {
		memPort = "11211"
	}

	mc := memcache.New(fmt.Sprintf("%s:%s", memHost, memPort))
	s := newState(mc)

	handleRequests(s)
}

func handleRequests(s *state) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get/{key}", s.handleGet)
	router.HandleFunc("/set/{key}/{value}", s.handleSet)

	log.Print("Starting on http://localhost:8989")
	if err := http.ListenAndServe(":8989", router); err != nil {
		log.Fatalf("cannot listen on port 8989 %v", err)
	}
}
