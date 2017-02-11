package main

import (
	"fmt"
	"net/http"
)

type httpHandler func(http.ResponseWriter, *http.Request)

func setContentheader(h httpHandler) httpHandler{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content=Type", "application/json;charset=UTF-8")
		h(w,r)
	}
}
func handler(w http.ResponseWriter, r *http.Request){
	body, err := sayer(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.Write(body)
}

func sayer( r *http.Request)([]byte, error){ 
	m := "I'm a test"
	o := []byte(m)
	return o, nil

}

func pants(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "i'm a pants test")
}

func main(){
	http.HandleFunc("/pants/", setContentheader(pants))
	http.HandleFunc("/", setContentheader(handler))


	http.ListenAndServe(":8080", nil)
}