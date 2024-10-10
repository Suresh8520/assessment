package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", FileHandler)

	fmt.Println("Server started, running 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(data))
}
