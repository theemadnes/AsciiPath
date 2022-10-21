package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// set default port number if env var $PORT isn't set
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		fmt.Printf("Req: %s\n", r.URL.Path)
		//http.NotFound(w, r)
		str := strings.Replace(r.URL.Path, "/", " ", -1) // replace slashes in path with spaces
		asciiStr := figure.NewFigure(str, "", true)
		io.WriteString(w, asciiStr.String())
		return
	}
	fmt.Printf("got / request\n")
	str := "base path"
	asciiStr := figure.NewFigure(str, "", true)
	io.WriteString(w, asciiStr.String())
}

func main() {
	port := getEnv("PORT", "8080")

	//http.HandleFunc("/*")
	http.HandleFunc("/", getHandler)

	http.ListenAndServe(":"+port, nil)
}
