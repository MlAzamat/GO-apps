package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	
	var i int
	i++
	//var s string = fmt.Sprintf("Hello GOLANG %d", i)

	fmt.Fprintf(w, "Hi there, I love People # %d ! Path /%s!\n", i, r.URL.Path[1:])	
	w.Write([]byte("HELLO"))
}

func main() {
	http.HandleFunc("/people", Handler)
	log.Println("Status HTTP server on port 8081")
	log.Fatal(http.ListenAndServe(":8081",nil))
}


