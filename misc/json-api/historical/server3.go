package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func fatCatHandler(w http.ResponseWriter, r *http.Request) {
	b, err := fatCatFacts(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("There was an error loading Cat Facts: %s", err)
		return
	}

	if _, err := w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("There was an error writing\n Error: %s", err)
	}
}
func fatCatFacts(r *http.Request) ([]byte, error) {
	m := map[string]string{
		"weight": "Claws can weigh up to 20,000 pounds!",
		"noise":  "Claws can be heard yodeling 4 miles away in good weather.",
		"cuddle": "Claws do not cuddle, they take affection.",
	}
	if val, ok := m[r.URL.Path[len("/api/fatcat/"):]]; ok {
		t := struct {
			Term string `json: "term"`
			Fact string `json: "fact"`
		}{
			Term: m[r.URL.Path[len("/api/fatcat/"):]],
			Fact: val,
		}

		return json.Marshal(t)
	}
	return []byte("Fat cat fact not found!"), nil

}

func main() {
	http.HandleFunc("/api/fatcat/", fatCatHandler)
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}
