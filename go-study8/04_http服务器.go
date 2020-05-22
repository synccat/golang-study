package main

import (
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("helllo world"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}

}
