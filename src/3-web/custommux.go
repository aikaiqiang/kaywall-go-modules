package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		//sayHelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	of := reflect.TypeOf(mux)
	fmt.Println(of)
	http.ListenAndServe(":8080", mux)
}
