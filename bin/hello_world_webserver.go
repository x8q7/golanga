package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

type obj struct {
	name string
}

func (o *obj) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("http input:  " + req.URL.Path)
	io.WriteString(w, "hello "+o.name)
}

func main() {
	http.HandleFunc("/", HelloServer)

	o1 := &obj{name: "Chris"}
	o2 := &obj{name: "Madeleine"}
	http.Handle("/hello/Name", o1)
	http.Handle("/shouthello/Name", o2)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
