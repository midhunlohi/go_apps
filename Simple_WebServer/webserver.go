package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int

var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {

	http.HandleFunc("/increment", incrementCounter)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil))

}
