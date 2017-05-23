package main

import (
	"characters"
	"log"
	"myHTTP"
	"net/http"
	"os"
)

const (
	unkown               = iota
	contentType          = "Content-Type"
	jsonContentType      = "application/json; charset=utf-8"
	plainTextContentType = "text/plain; charset=utf-8"
)

type charComp func(r *http.Request) ([]byte, error)

func magicHandler(f charComp) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpCode := http.StatusInternalServerError

		b, err := f(r)
		if e, ok := err.(*myHTTP.HTTPError); ok {
			httpCode = e.GetCode()
			if e != myHTTP.OK {
				w.Header().Set(contentType, plainTextContentType)
				w.WriteHeader(httpCode)
				w.Write([]byte(e.Err))
				return
			}
		}
		if _, err := w.Write(b); err != nil {
			w.WriteHeader(httpCode)
			return
		}
		w.Header().Set(contentType, jsonContentType)
		w.WriteHeader(httpCode)
	}
}
func main() {
	//	http.HandleFunc("/api/spells/", magicHandler(spells.Dispatcher))
	http.HandleFunc("/api/characters/", magicHandler(characters.Dispatcher))
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
