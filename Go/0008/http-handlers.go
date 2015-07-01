package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

type Hello struct{}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func main() {
	//http.Handle("/", Hello{})
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	//err := http.ListenAndServe("localhost:4000", nil)

	var h Hello
	err := http.ListenAndServe("localhost:4000", h)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
