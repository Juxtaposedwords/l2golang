package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type apiFunc func(r *http.Request) (*apiResponse, error)
type apiResponse struct {
	Name string `json:"name"`
}

func wrapApiFunc(f apiFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ar, err := f(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(*ar)
		if err != nil {
			log.Printf("error marshalling %v: %s", ar, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(b); err != nil {
			log.Printf("Error writing response: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func apiFunc1(r *http.Request) (*apiResponse, error) {
	return &apiResponse{"Calling func1"}, nil
}

func apiFunc2(r *http.Request) (*apiResponse, error) {
	return &apiResponse{"Calling func2"}, nil
}

func apiFunc3(r *http.Request) (*apiResponse, error) {
	return nil, errors.New("apiFunc3 always returns an error")
}

func main() {
	http.HandleFunc("/api/func1", wrapApiFunc(apiFunc1))
	http.HandleFunc("/api/func2", wrapApiFunc(apiFunc2))
	http.HandleFunc("/api/func3", wrapApiFunc(apiFunc3))
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
