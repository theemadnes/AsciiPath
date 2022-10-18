package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// set default port number if env var $PORT isn't set
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	port := getEnv("PORT", "8080")

	http.HandleFunc("/", getRoot)

	http.ListenAndServe(":"+port, nil)
}
